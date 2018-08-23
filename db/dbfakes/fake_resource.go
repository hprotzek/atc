// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

type FakeResource struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PipelineNameStub        func() string
	pipelineNameMutex       sync.RWMutex
	pipelineNameArgsForCall []struct{}
	pipelineNameReturns     struct {
		result1 string
	}
	pipelineNameReturnsOnCall map[int]struct {
		result1 string
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct{}
	teamNameReturns     struct {
		result1 string
	}
	teamNameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct{}
	sourceReturns     struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	CheckEveryStub        func() string
	checkEveryMutex       sync.RWMutex
	checkEveryArgsForCall []struct{}
	checkEveryReturns     struct {
		result1 string
	}
	checkEveryReturnsOnCall map[int]struct {
		result1 string
	}
	CheckTimeoutStub        func() string
	checkTimeoutMutex       sync.RWMutex
	checkTimeoutArgsForCall []struct{}
	checkTimeoutReturns     struct {
		result1 string
	}
	checkTimeoutReturnsOnCall map[int]struct {
		result1 string
	}
	LastCheckedStub        func() time.Time
	lastCheckedMutex       sync.RWMutex
	lastCheckedArgsForCall []struct{}
	lastCheckedReturns     struct {
		result1 time.Time
	}
	lastCheckedReturnsOnCall map[int]struct {
		result1 time.Time
	}
	TagsStub        func() atc.Tags
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct{}
	tagsReturns     struct {
		result1 atc.Tags
	}
	tagsReturnsOnCall map[int]struct {
		result1 atc.Tags
	}
	CheckErrorStub        func() error
	checkErrorMutex       sync.RWMutex
	checkErrorArgsForCall []struct{}
	checkErrorReturns     struct {
		result1 error
	}
	checkErrorReturnsOnCall map[int]struct {
		result1 error
	}
	PausedStub        func() bool
	pausedMutex       sync.RWMutex
	pausedArgsForCall []struct{}
	pausedReturns     struct {
		result1 bool
	}
	pausedReturnsOnCall map[int]struct {
		result1 bool
	}
	WebhookTokenStub        func() string
	webhookTokenMutex       sync.RWMutex
	webhookTokenArgsForCall []struct{}
	webhookTokenReturns     struct {
		result1 string
	}
	webhookTokenReturnsOnCall map[int]struct {
		result1 string
	}
	PinnedVersionStub        func() atc.Version
	pinnedVersionMutex       sync.RWMutex
	pinnedVersionArgsForCall []struct{}
	pinnedVersionReturns     struct {
		result1 atc.Version
	}
	pinnedVersionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	FailingToCheckStub        func() bool
	failingToCheckMutex       sync.RWMutex
	failingToCheckArgsForCall []struct{}
	failingToCheckReturns     struct {
		result1 bool
	}
	failingToCheckReturnsOnCall map[int]struct {
		result1 bool
	}
	ResourceConfigCheckErrorStub        func() error
	resourceConfigCheckErrorMutex       sync.RWMutex
	resourceConfigCheckErrorArgsForCall []struct{}
	resourceConfigCheckErrorReturns     struct {
		result1 error
	}
	resourceConfigCheckErrorReturnsOnCall map[int]struct {
		result1 error
	}
	SetResourceConfigStub        func(int) error
	setResourceConfigMutex       sync.RWMutex
	setResourceConfigArgsForCall []struct {
		arg1 int
	}
	setResourceConfigReturns struct {
		result1 error
	}
	setResourceConfigReturnsOnCall map[int]struct {
		result1 error
	}
	SetCheckErrorStub        func(error) error
	setCheckErrorMutex       sync.RWMutex
	setCheckErrorArgsForCall []struct {
		arg1 error
	}
	setCheckErrorReturns struct {
		result1 error
	}
	setCheckErrorReturnsOnCall map[int]struct {
		result1 error
	}
	PauseStub        func() error
	pauseMutex       sync.RWMutex
	pauseArgsForCall []struct{}
	pauseReturns     struct {
		result1 error
	}
	pauseReturnsOnCall map[int]struct {
		result1 error
	}
	UnpauseStub        func() error
	unpauseMutex       sync.RWMutex
	unpauseArgsForCall []struct{}
	unpauseReturns     struct {
		result1 error
	}
	unpauseReturnsOnCall map[int]struct {
		result1 error
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResource) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeResource) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResource) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResource) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResource) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeResource) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeResource) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) PipelineName() string {
	fake.pipelineNameMutex.Lock()
	ret, specificReturn := fake.pipelineNameReturnsOnCall[len(fake.pipelineNameArgsForCall)]
	fake.pipelineNameArgsForCall = append(fake.pipelineNameArgsForCall, struct{}{})
	fake.recordInvocation("PipelineName", []interface{}{})
	fake.pipelineNameMutex.Unlock()
	if fake.PipelineNameStub != nil {
		return fake.PipelineNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pipelineNameReturns.result1
}

func (fake *FakeResource) PipelineNameCallCount() int {
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	return len(fake.pipelineNameArgsForCall)
}

func (fake *FakeResource) PipelineNameReturns(result1 string) {
	fake.PipelineNameStub = nil
	fake.pipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) PipelineNameReturnsOnCall(i int, result1 string) {
	fake.PipelineNameStub = nil
	if fake.pipelineNameReturnsOnCall == nil {
		fake.pipelineNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pipelineNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) TeamName() string {
	fake.teamNameMutex.Lock()
	ret, specificReturn := fake.teamNameReturnsOnCall[len(fake.teamNameArgsForCall)]
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct{}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.teamNameReturns.result1
}

func (fake *FakeResource) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeResource) TeamNameReturns(result1 string) {
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) TeamNameReturnsOnCall(i int, result1 string) {
	fake.TeamNameStub = nil
	if fake.teamNameReturnsOnCall == nil {
		fake.teamNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.teamNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeResource) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeResource) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) TypeReturnsOnCall(i int, result1 string) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct{}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sourceReturns.result1
}

func (fake *FakeResource) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeResource) SourceReturns(result1 atc.Source) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResource) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResource) CheckEvery() string {
	fake.checkEveryMutex.Lock()
	ret, specificReturn := fake.checkEveryReturnsOnCall[len(fake.checkEveryArgsForCall)]
	fake.checkEveryArgsForCall = append(fake.checkEveryArgsForCall, struct{}{})
	fake.recordInvocation("CheckEvery", []interface{}{})
	fake.checkEveryMutex.Unlock()
	if fake.CheckEveryStub != nil {
		return fake.CheckEveryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkEveryReturns.result1
}

func (fake *FakeResource) CheckEveryCallCount() int {
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	return len(fake.checkEveryArgsForCall)
}

func (fake *FakeResource) CheckEveryReturns(result1 string) {
	fake.CheckEveryStub = nil
	fake.checkEveryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) CheckEveryReturnsOnCall(i int, result1 string) {
	fake.CheckEveryStub = nil
	if fake.checkEveryReturnsOnCall == nil {
		fake.checkEveryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkEveryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) CheckTimeout() string {
	fake.checkTimeoutMutex.Lock()
	ret, specificReturn := fake.checkTimeoutReturnsOnCall[len(fake.checkTimeoutArgsForCall)]
	fake.checkTimeoutArgsForCall = append(fake.checkTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("CheckTimeout", []interface{}{})
	fake.checkTimeoutMutex.Unlock()
	if fake.CheckTimeoutStub != nil {
		return fake.CheckTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkTimeoutReturns.result1
}

func (fake *FakeResource) CheckTimeoutCallCount() int {
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	return len(fake.checkTimeoutArgsForCall)
}

func (fake *FakeResource) CheckTimeoutReturns(result1 string) {
	fake.CheckTimeoutStub = nil
	fake.checkTimeoutReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) CheckTimeoutReturnsOnCall(i int, result1 string) {
	fake.CheckTimeoutStub = nil
	if fake.checkTimeoutReturnsOnCall == nil {
		fake.checkTimeoutReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkTimeoutReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) LastChecked() time.Time {
	fake.lastCheckedMutex.Lock()
	ret, specificReturn := fake.lastCheckedReturnsOnCall[len(fake.lastCheckedArgsForCall)]
	fake.lastCheckedArgsForCall = append(fake.lastCheckedArgsForCall, struct{}{})
	fake.recordInvocation("LastChecked", []interface{}{})
	fake.lastCheckedMutex.Unlock()
	if fake.LastCheckedStub != nil {
		return fake.LastCheckedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lastCheckedReturns.result1
}

func (fake *FakeResource) LastCheckedCallCount() int {
	fake.lastCheckedMutex.RLock()
	defer fake.lastCheckedMutex.RUnlock()
	return len(fake.lastCheckedArgsForCall)
}

func (fake *FakeResource) LastCheckedReturns(result1 time.Time) {
	fake.LastCheckedStub = nil
	fake.lastCheckedReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeResource) LastCheckedReturnsOnCall(i int, result1 time.Time) {
	fake.LastCheckedStub = nil
	if fake.lastCheckedReturnsOnCall == nil {
		fake.lastCheckedReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.lastCheckedReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeResource) Tags() atc.Tags {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct{}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tagsReturns.result1
}

func (fake *FakeResource) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeResource) TagsReturns(result1 atc.Tags) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeResource) TagsReturnsOnCall(i int, result1 atc.Tags) {
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 atc.Tags
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeResource) CheckError() error {
	fake.checkErrorMutex.Lock()
	ret, specificReturn := fake.checkErrorReturnsOnCall[len(fake.checkErrorArgsForCall)]
	fake.checkErrorArgsForCall = append(fake.checkErrorArgsForCall, struct{}{})
	fake.recordInvocation("CheckError", []interface{}{})
	fake.checkErrorMutex.Unlock()
	if fake.CheckErrorStub != nil {
		return fake.CheckErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkErrorReturns.result1
}

func (fake *FakeResource) CheckErrorCallCount() int {
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	return len(fake.checkErrorArgsForCall)
}

func (fake *FakeResource) CheckErrorReturns(result1 error) {
	fake.CheckErrorStub = nil
	fake.checkErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) CheckErrorReturnsOnCall(i int, result1 error) {
	fake.CheckErrorStub = nil
	if fake.checkErrorReturnsOnCall == nil {
		fake.checkErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) Paused() bool {
	fake.pausedMutex.Lock()
	ret, specificReturn := fake.pausedReturnsOnCall[len(fake.pausedArgsForCall)]
	fake.pausedArgsForCall = append(fake.pausedArgsForCall, struct{}{})
	fake.recordInvocation("Paused", []interface{}{})
	fake.pausedMutex.Unlock()
	if fake.PausedStub != nil {
		return fake.PausedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pausedReturns.result1
}

func (fake *FakeResource) PausedCallCount() int {
	fake.pausedMutex.RLock()
	defer fake.pausedMutex.RUnlock()
	return len(fake.pausedArgsForCall)
}

func (fake *FakeResource) PausedReturns(result1 bool) {
	fake.PausedStub = nil
	fake.pausedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResource) PausedReturnsOnCall(i int, result1 bool) {
	fake.PausedStub = nil
	if fake.pausedReturnsOnCall == nil {
		fake.pausedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.pausedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResource) WebhookToken() string {
	fake.webhookTokenMutex.Lock()
	ret, specificReturn := fake.webhookTokenReturnsOnCall[len(fake.webhookTokenArgsForCall)]
	fake.webhookTokenArgsForCall = append(fake.webhookTokenArgsForCall, struct{}{})
	fake.recordInvocation("WebhookToken", []interface{}{})
	fake.webhookTokenMutex.Unlock()
	if fake.WebhookTokenStub != nil {
		return fake.WebhookTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.webhookTokenReturns.result1
}

func (fake *FakeResource) WebhookTokenCallCount() int {
	fake.webhookTokenMutex.RLock()
	defer fake.webhookTokenMutex.RUnlock()
	return len(fake.webhookTokenArgsForCall)
}

func (fake *FakeResource) WebhookTokenReturns(result1 string) {
	fake.WebhookTokenStub = nil
	fake.webhookTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) WebhookTokenReturnsOnCall(i int, result1 string) {
	fake.WebhookTokenStub = nil
	if fake.webhookTokenReturnsOnCall == nil {
		fake.webhookTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.webhookTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) PinnedVersion() atc.Version {
	fake.pinnedVersionMutex.Lock()
	ret, specificReturn := fake.pinnedVersionReturnsOnCall[len(fake.pinnedVersionArgsForCall)]
	fake.pinnedVersionArgsForCall = append(fake.pinnedVersionArgsForCall, struct{}{})
	fake.recordInvocation("PinnedVersion", []interface{}{})
	fake.pinnedVersionMutex.Unlock()
	if fake.PinnedVersionStub != nil {
		return fake.PinnedVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pinnedVersionReturns.result1
}

func (fake *FakeResource) PinnedVersionCallCount() int {
	fake.pinnedVersionMutex.RLock()
	defer fake.pinnedVersionMutex.RUnlock()
	return len(fake.pinnedVersionArgsForCall)
}

func (fake *FakeResource) PinnedVersionReturns(result1 atc.Version) {
	fake.PinnedVersionStub = nil
	fake.pinnedVersionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResource) PinnedVersionReturnsOnCall(i int, result1 atc.Version) {
	fake.PinnedVersionStub = nil
	if fake.pinnedVersionReturnsOnCall == nil {
		fake.pinnedVersionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.pinnedVersionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResource) FailingToCheck() bool {
	fake.failingToCheckMutex.Lock()
	ret, specificReturn := fake.failingToCheckReturnsOnCall[len(fake.failingToCheckArgsForCall)]
	fake.failingToCheckArgsForCall = append(fake.failingToCheckArgsForCall, struct{}{})
	fake.recordInvocation("FailingToCheck", []interface{}{})
	fake.failingToCheckMutex.Unlock()
	if fake.FailingToCheckStub != nil {
		return fake.FailingToCheckStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.failingToCheckReturns.result1
}

func (fake *FakeResource) FailingToCheckCallCount() int {
	fake.failingToCheckMutex.RLock()
	defer fake.failingToCheckMutex.RUnlock()
	return len(fake.failingToCheckArgsForCall)
}

func (fake *FakeResource) FailingToCheckReturns(result1 bool) {
	fake.FailingToCheckStub = nil
	fake.failingToCheckReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResource) FailingToCheckReturnsOnCall(i int, result1 bool) {
	fake.FailingToCheckStub = nil
	if fake.failingToCheckReturnsOnCall == nil {
		fake.failingToCheckReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.failingToCheckReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResource) ResourceConfigCheckError() error {
	fake.resourceConfigCheckErrorMutex.Lock()
	ret, specificReturn := fake.resourceConfigCheckErrorReturnsOnCall[len(fake.resourceConfigCheckErrorArgsForCall)]
	fake.resourceConfigCheckErrorArgsForCall = append(fake.resourceConfigCheckErrorArgsForCall, struct{}{})
	fake.recordInvocation("ResourceConfigCheckError", []interface{}{})
	fake.resourceConfigCheckErrorMutex.Unlock()
	if fake.ResourceConfigCheckErrorStub != nil {
		return fake.ResourceConfigCheckErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceConfigCheckErrorReturns.result1
}

func (fake *FakeResource) ResourceConfigCheckErrorCallCount() int {
	fake.resourceConfigCheckErrorMutex.RLock()
	defer fake.resourceConfigCheckErrorMutex.RUnlock()
	return len(fake.resourceConfigCheckErrorArgsForCall)
}

func (fake *FakeResource) ResourceConfigCheckErrorReturns(result1 error) {
	fake.ResourceConfigCheckErrorStub = nil
	fake.resourceConfigCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) ResourceConfigCheckErrorReturnsOnCall(i int, result1 error) {
	fake.ResourceConfigCheckErrorStub = nil
	if fake.resourceConfigCheckErrorReturnsOnCall == nil {
		fake.resourceConfigCheckErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resourceConfigCheckErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) SetResourceConfig(arg1 int) error {
	fake.setResourceConfigMutex.Lock()
	ret, specificReturn := fake.setResourceConfigReturnsOnCall[len(fake.setResourceConfigArgsForCall)]
	fake.setResourceConfigArgsForCall = append(fake.setResourceConfigArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetResourceConfig", []interface{}{arg1})
	fake.setResourceConfigMutex.Unlock()
	if fake.SetResourceConfigStub != nil {
		return fake.SetResourceConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setResourceConfigReturns.result1
}

func (fake *FakeResource) SetResourceConfigCallCount() int {
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	return len(fake.setResourceConfigArgsForCall)
}

func (fake *FakeResource) SetResourceConfigArgsForCall(i int) int {
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	return fake.setResourceConfigArgsForCall[i].arg1
}

func (fake *FakeResource) SetResourceConfigReturns(result1 error) {
	fake.SetResourceConfigStub = nil
	fake.setResourceConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) SetResourceConfigReturnsOnCall(i int, result1 error) {
	fake.SetResourceConfigStub = nil
	if fake.setResourceConfigReturnsOnCall == nil {
		fake.setResourceConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setResourceConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) SetCheckError(arg1 error) error {
	fake.setCheckErrorMutex.Lock()
	ret, specificReturn := fake.setCheckErrorReturnsOnCall[len(fake.setCheckErrorArgsForCall)]
	fake.setCheckErrorArgsForCall = append(fake.setCheckErrorArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.recordInvocation("SetCheckError", []interface{}{arg1})
	fake.setCheckErrorMutex.Unlock()
	if fake.SetCheckErrorStub != nil {
		return fake.SetCheckErrorStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setCheckErrorReturns.result1
}

func (fake *FakeResource) SetCheckErrorCallCount() int {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	return len(fake.setCheckErrorArgsForCall)
}

func (fake *FakeResource) SetCheckErrorArgsForCall(i int) error {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	return fake.setCheckErrorArgsForCall[i].arg1
}

func (fake *FakeResource) SetCheckErrorReturns(result1 error) {
	fake.SetCheckErrorStub = nil
	fake.setCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) SetCheckErrorReturnsOnCall(i int, result1 error) {
	fake.SetCheckErrorStub = nil
	if fake.setCheckErrorReturnsOnCall == nil {
		fake.setCheckErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setCheckErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) Pause() error {
	fake.pauseMutex.Lock()
	ret, specificReturn := fake.pauseReturnsOnCall[len(fake.pauseArgsForCall)]
	fake.pauseArgsForCall = append(fake.pauseArgsForCall, struct{}{})
	fake.recordInvocation("Pause", []interface{}{})
	fake.pauseMutex.Unlock()
	if fake.PauseStub != nil {
		return fake.PauseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pauseReturns.result1
}

func (fake *FakeResource) PauseCallCount() int {
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	return len(fake.pauseArgsForCall)
}

func (fake *FakeResource) PauseReturns(result1 error) {
	fake.PauseStub = nil
	fake.pauseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) PauseReturnsOnCall(i int, result1 error) {
	fake.PauseStub = nil
	if fake.pauseReturnsOnCall == nil {
		fake.pauseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pauseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) Unpause() error {
	fake.unpauseMutex.Lock()
	ret, specificReturn := fake.unpauseReturnsOnCall[len(fake.unpauseArgsForCall)]
	fake.unpauseArgsForCall = append(fake.unpauseArgsForCall, struct{}{})
	fake.recordInvocation("Unpause", []interface{}{})
	fake.unpauseMutex.Unlock()
	if fake.UnpauseStub != nil {
		return fake.UnpauseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unpauseReturns.result1
}

func (fake *FakeResource) UnpauseCallCount() int {
	fake.unpauseMutex.RLock()
	defer fake.unpauseMutex.RUnlock()
	return len(fake.unpauseArgsForCall)
}

func (fake *FakeResource) UnpauseReturns(result1 error) {
	fake.UnpauseStub = nil
	fake.unpauseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) UnpauseReturnsOnCall(i int, result1 error) {
	fake.UnpauseStub = nil
	if fake.unpauseReturnsOnCall == nil {
		fake.unpauseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unpauseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.reloadReturns.result1, fake.reloadReturns.result2
}

func (fake *FakeResource) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeResource) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	fake.lastCheckedMutex.RLock()
	defer fake.lastCheckedMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	fake.pausedMutex.RLock()
	defer fake.pausedMutex.RUnlock()
	fake.webhookTokenMutex.RLock()
	defer fake.webhookTokenMutex.RUnlock()
	fake.pinnedVersionMutex.RLock()
	defer fake.pinnedVersionMutex.RUnlock()
	fake.failingToCheckMutex.RLock()
	defer fake.failingToCheckMutex.RUnlock()
	fake.resourceConfigCheckErrorMutex.RLock()
	defer fake.resourceConfigCheckErrorMutex.RUnlock()
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	fake.unpauseMutex.RLock()
	defer fake.unpauseMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Resource = new(FakeResource)
