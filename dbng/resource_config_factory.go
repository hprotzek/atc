package dbng

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/atc"
)

//go:generate counterfeiter . ResourceConfigFactory

type ResourceConfigFactory interface {
	FindOrCreateResourceConfigForBuild(
		build *Build,
		resourceType string,
		source atc.Source,
		pipeline *Pipeline,
		resourceTypes atc.ResourceTypes,
	) (*UsedResourceConfig, error)

	FindOrCreateResourceConfigForResource(
		resource *Resource,
		resourceType string,
		source atc.Source,
		pipeline *Pipeline,
		resourceTypes atc.ResourceTypes,
	) (*UsedResourceConfig, error)

	FindOrCreateResourceConfigForResourceType(
		resourceTypeName string,
		source atc.Source,
		pipeline *Pipeline,
		resourceTypes atc.ResourceTypes,
	) (*UsedResourceConfig, error)

	CleanConfigUsesForFinishedBuilds() error
	CleanConfigUsesForInactiveResourceTypes() error
	CleanConfigUsesForInactiveResources() error
	CleanUselessConfigs() error
}

type resourceConfigFactory struct {
	conn Conn
}

func NewResourceConfigFactory(conn Conn) ResourceConfigFactory {
	return &resourceConfigFactory{
		conn: conn,
	}
}

func (f *resourceConfigFactory) FindOrCreateResourceConfigForBuild(
	build *Build,
	resourceType string,
	source atc.Source,
	pipeline *Pipeline,
	resourceTypes atc.ResourceTypes,
) (*UsedResourceConfig, error) {
	tx, err := f.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbResourceTypes, err := getDBResourceTypes(tx, pipeline, resourceTypes)
	if err != nil {
		return nil, err
	}

	resourceConfig, err := constructResourceConfig(resourceType, source, dbResourceTypes)
	if err != nil {
		return nil, err
	}

	usedResourceConfig, err := resourceConfig.FindOrCreateForBuild(tx, build)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return usedResourceConfig, nil
}

func (f *resourceConfigFactory) FindOrCreateResourceConfigForResource(
	resource *Resource,
	resourceType string,
	source atc.Source,
	pipeline *Pipeline,
	resourceTypes atc.ResourceTypes,
) (*UsedResourceConfig, error) {
	tx, err := f.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbResourceTypes, err := getDBResourceTypes(tx, pipeline, resourceTypes)
	if err != nil {
		return nil, err
	}

	resourceConfig, err := constructResourceConfig(resourceType, source, dbResourceTypes)
	if err != nil {
		return nil, err
	}

	usedResourceConfig, err := resourceConfig.FindOrCreateForResource(tx, resource)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return usedResourceConfig, nil
}

func (f *resourceConfigFactory) FindOrCreateResourceConfigForResourceType(
	resourceTypeName string,
	source atc.Source,
	pipeline *Pipeline,
	resourceTypes atc.ResourceTypes,
) (*UsedResourceConfig, error) {
	resourceType, found := resourceTypes.Lookup(resourceTypeName)
	if !found {
		return nil, ErrResourceTypeNotFound{resourceTypeName}
	}

	tx, err := f.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbResourceTypes, err := getDBResourceTypes(tx, pipeline, resourceTypes)
	if err != nil {
		return nil, err
	}

	resourceConfig, err := constructResourceConfig(resourceTypeName, source, dbResourceTypes)
	if err != nil {
		return nil, err
	}

	rt := ResourceType{
		ResourceType: resourceType,
		Pipeline:     pipeline,
	}

	usedResourceType, found, err := rt.Find(tx)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, ErrResourceTypeNotFound{resourceTypeName}
	}

	usedResourceConfig, err := resourceConfig.FindOrCreateForResourceType(tx, usedResourceType)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return usedResourceConfig, nil
}

func constructResourceConfig(
	resourceType string,
	source atc.Source,
	resourceTypes []ResourceType,
) (ResourceConfig, error) {
	resourceConfig := ResourceConfig{
		Source: source,
	}

	resourceTypesList := resourceTypesList(resourceType, resourceTypes, []ResourceType{})
	if len(resourceTypesList) == 0 {
		resourceConfig.CreatedByBaseResourceType = &BaseResourceType{
			Name: resourceType,
		}
	} else {
		lastResourceType := resourceTypesList[len(resourceTypesList)-1]

		parentResourceCache := &ResourceCache{
			ResourceConfig: ResourceConfig{
				CreatedByBaseResourceType: &BaseResourceType{
					Name: lastResourceType.Type,
				},
				Source: lastResourceType.Source,
			},
			Version: lastResourceType.Version,
		}

		for i := len(resourceTypesList) - 2; i >= 0; i-- {
			parentResourceCache = &ResourceCache{
				ResourceConfig: ResourceConfig{
					CreatedByResourceCache: parentResourceCache,
					Source:                 resourceTypesList[i].Source,
				},
				Version: resourceTypesList[i].Version,
			}
		}

		resourceConfig.CreatedByResourceCache = parentResourceCache
	}

	return resourceConfig, nil
}

func (f *resourceConfigFactory) CleanConfigUsesForFinishedBuilds() error {
	tx, err := f.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	latestBuildByJobQ, _, err := sq.
		Select("MAX(b.id) AS build_id", "j.id AS job_id").
		From("builds b").
		Join("jobs j ON j.id = b.job_id").
		GroupBy("j.id").ToSql()
	if err != nil {
		return err
	}

	extractedBuildIds, _, err := sq.
		Select("lbbjq.build_id").
		Distinct().
		From("(" + latestBuildByJobQ + ") as lbbjq").
		ToSql()
	if err != nil {
		return err
	}

	_, err = psql.Delete("resource_config_uses rcu USING builds b").
		Where(sq.And{
			sq.Expr("rcu.build_id = b.id"),
			sq.Or{
				sq.Eq{
					"b.status": "succeeded",
				},
				sq.And{
					sq.Expr("b.id NOT IN (" + extractedBuildIds + ")"),
					sq.Eq{
						"b.status": "failed",
					},
				},
				sq.Eq{
					"b.status": "aborted",
				},
			},
		}).
		RunWith(tx).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (f *resourceConfigFactory) CleanConfigUsesForInactiveResourceTypes() error {
	tx, err := f.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = psql.Delete("resource_config_uses rcu USING resource_types t").
		Where(sq.And{
			sq.Expr("rcu.resource_type_id = t.id"),
			sq.Eq{
				"t.active": false,
			},
		}).
		RunWith(tx).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (f *resourceConfigFactory) CleanConfigUsesForInactiveResources() error {
	tx, err := f.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = psql.Delete("resource_config_uses rcu USING resources r").
		Where(sq.And{
			sq.Expr("rcu.resource_id = r.id"),
			sq.Eq{
				"r.active": false,
			},
		}).
		RunWith(tx).
		Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (f *resourceConfigFactory) CleanUselessConfigs() error {
	tx, err := f.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stillInUseConfigIds, _, err := sq.
		Select("rf.id").
		Distinct().
		From("resource_configs rf").
		JoinClause("INNER JOIN resource_config_uses rfu ON rf.id = rfu.resource_config_id").
		ToSql()
	if err != nil {
		return err
	}

	_, err = psql.Delete("resource_configs").
		Where("id NOT IN (" + stillInUseConfigIds + ")").
		PlaceholderFormat(sq.Dollar).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func resourceTypesList(resourceTypeName string, allResourceTypes []ResourceType, resultResourceTypes []ResourceType) []ResourceType {
	for _, resourceType := range allResourceTypes {
		if resourceType.Name == resourceTypeName {
			resultResourceTypes = append(resultResourceTypes, resourceType)
			return resourceTypesList(resourceType.Type, allResourceTypes, resultResourceTypes)
		}
	}

	return resultResourceTypes
}