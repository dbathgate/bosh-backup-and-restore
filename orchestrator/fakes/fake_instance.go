// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/bosh-backup-and-restore/orchestrator"
)

type FakeInstance struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	IndexStub        func() string
	indexMutex       sync.RWMutex
	indexArgsForCall []struct{}
	indexReturns     struct {
		result1 string
	}
	indexReturnsOnCall map[int]struct {
		result1 string
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	IsBackupableStub        func() bool
	isBackupableMutex       sync.RWMutex
	isBackupableArgsForCall []struct{}
	isBackupableReturns     struct {
		result1 bool
	}
	isBackupableReturnsOnCall map[int]struct {
		result1 bool
	}
	ArtifactDirExistsStub        func() (bool, error)
	artifactDirExistsMutex       sync.RWMutex
	artifactDirExistsArgsForCall []struct{}
	artifactDirExistsReturns     struct {
		result1 bool
		result2 error
	}
	artifactDirExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ArtifactDirCreatedStub        func() bool
	artifactDirCreatedMutex       sync.RWMutex
	artifactDirCreatedArgsForCall []struct{}
	artifactDirCreatedReturns     struct {
		result1 bool
	}
	artifactDirCreatedReturnsOnCall map[int]struct {
		result1 bool
	}
	MarkArtifactDirCreatedStub        func()
	markArtifactDirCreatedMutex       sync.RWMutex
	markArtifactDirCreatedArgsForCall []struct{}
	IsPostBackupUnlockableStub        func() bool
	isPostBackupUnlockableMutex       sync.RWMutex
	isPostBackupUnlockableArgsForCall []struct{}
	isPostBackupUnlockableReturns     struct {
		result1 bool
	}
	isPostBackupUnlockableReturnsOnCall map[int]struct {
		result1 bool
	}
	IsPreBackupLockableStub        func() bool
	isPreBackupLockableMutex       sync.RWMutex
	isPreBackupLockableArgsForCall []struct{}
	isPreBackupLockableReturns     struct {
		result1 bool
	}
	isPreBackupLockableReturnsOnCall map[int]struct {
		result1 bool
	}
	IsRestorableStub        func() bool
	isRestorableMutex       sync.RWMutex
	isRestorableArgsForCall []struct{}
	isRestorableReturns     struct {
		result1 bool
	}
	isRestorableReturnsOnCall map[int]struct {
		result1 bool
	}
	PreBackupLockStub        func() error
	preBackupLockMutex       sync.RWMutex
	preBackupLockArgsForCall []struct{}
	preBackupLockReturns     struct {
		result1 error
	}
	preBackupLockReturnsOnCall map[int]struct {
		result1 error
	}
	BackupStub        func() error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct{}
	backupReturns     struct {
		result1 error
	}
	backupReturnsOnCall map[int]struct {
		result1 error
	}
	PostBackupUnlockStub        func() error
	postBackupUnlockMutex       sync.RWMutex
	postBackupUnlockArgsForCall []struct{}
	postBackupUnlockReturns     struct {
		result1 error
	}
	postBackupUnlockReturnsOnCall map[int]struct {
		result1 error
	}
	RestoreStub        func() error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct{}
	restoreReturns     struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	CleanupPreviousStub        func() error
	cleanupPreviousMutex       sync.RWMutex
	cleanupPreviousArgsForCall []struct{}
	cleanupPreviousReturns     struct {
		result1 error
	}
	cleanupPreviousReturnsOnCall map[int]struct {
		result1 error
	}
	ArtifactsToBackupStub        func() []orchestrator.BackupArtifact
	artifactsToBackupMutex       sync.RWMutex
	artifactsToBackupArgsForCall []struct{}
	artifactsToBackupReturns     struct {
		result1 []orchestrator.BackupArtifact
	}
	artifactsToBackupReturnsOnCall map[int]struct {
		result1 []orchestrator.BackupArtifact
	}
	ArtifactsToRestoreStub        func() []orchestrator.BackupArtifact
	artifactsToRestoreMutex       sync.RWMutex
	artifactsToRestoreArgsForCall []struct{}
	artifactsToRestoreReturns     struct {
		result1 []orchestrator.BackupArtifact
	}
	artifactsToRestoreReturnsOnCall map[int]struct {
		result1 []orchestrator.BackupArtifact
	}
	CustomBackupArtifactNamesStub        func() []string
	customBackupArtifactNamesMutex       sync.RWMutex
	customBackupArtifactNamesArgsForCall []struct{}
	customBackupArtifactNamesReturns     struct {
		result1 []string
	}
	customBackupArtifactNamesReturnsOnCall map[int]struct {
		result1 []string
	}
	CustomRestoreArtifactNamesStub        func() []string
	customRestoreArtifactNamesMutex       sync.RWMutex
	customRestoreArtifactNamesArgsForCall []struct{}
	customRestoreArtifactNamesReturns     struct {
		result1 []string
	}
	customRestoreArtifactNamesReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstance) Name() string {
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

func (fake *FakeInstance) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeInstance) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) NameReturnsOnCall(i int, result1 string) {
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

func (fake *FakeInstance) Index() string {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct{}{})
	fake.recordInvocation("Index", []interface{}{})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.indexReturns.result1
}

func (fake *FakeInstance) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeInstance) IndexReturns(result1 string) {
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) IndexReturnsOnCall(i int, result1 string) {
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) ID() string {
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

func (fake *FakeInstance) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeInstance) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) IDReturnsOnCall(i int, result1 string) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) IsBackupable() bool {
	fake.isBackupableMutex.Lock()
	ret, specificReturn := fake.isBackupableReturnsOnCall[len(fake.isBackupableArgsForCall)]
	fake.isBackupableArgsForCall = append(fake.isBackupableArgsForCall, struct{}{})
	fake.recordInvocation("IsBackupable", []interface{}{})
	fake.isBackupableMutex.Unlock()
	if fake.IsBackupableStub != nil {
		return fake.IsBackupableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isBackupableReturns.result1
}

func (fake *FakeInstance) IsBackupableCallCount() int {
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	return len(fake.isBackupableArgsForCall)
}

func (fake *FakeInstance) IsBackupableReturns(result1 bool) {
	fake.IsBackupableStub = nil
	fake.isBackupableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsBackupableReturnsOnCall(i int, result1 bool) {
	fake.IsBackupableStub = nil
	if fake.isBackupableReturnsOnCall == nil {
		fake.isBackupableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isBackupableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) ArtifactDirExists() (bool, error) {
	fake.artifactDirExistsMutex.Lock()
	ret, specificReturn := fake.artifactDirExistsReturnsOnCall[len(fake.artifactDirExistsArgsForCall)]
	fake.artifactDirExistsArgsForCall = append(fake.artifactDirExistsArgsForCall, struct{}{})
	fake.recordInvocation("ArtifactDirExists", []interface{}{})
	fake.artifactDirExistsMutex.Unlock()
	if fake.ArtifactDirExistsStub != nil {
		return fake.ArtifactDirExistsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.artifactDirExistsReturns.result1, fake.artifactDirExistsReturns.result2
}

func (fake *FakeInstance) ArtifactDirExistsCallCount() int {
	fake.artifactDirExistsMutex.RLock()
	defer fake.artifactDirExistsMutex.RUnlock()
	return len(fake.artifactDirExistsArgsForCall)
}

func (fake *FakeInstance) ArtifactDirExistsReturns(result1 bool, result2 error) {
	fake.ArtifactDirExistsStub = nil
	fake.artifactDirExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) ArtifactDirExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ArtifactDirExistsStub = nil
	if fake.artifactDirExistsReturnsOnCall == nil {
		fake.artifactDirExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.artifactDirExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) ArtifactDirCreated() bool {
	fake.artifactDirCreatedMutex.Lock()
	ret, specificReturn := fake.artifactDirCreatedReturnsOnCall[len(fake.artifactDirCreatedArgsForCall)]
	fake.artifactDirCreatedArgsForCall = append(fake.artifactDirCreatedArgsForCall, struct{}{})
	fake.recordInvocation("ArtifactDirCreated", []interface{}{})
	fake.artifactDirCreatedMutex.Unlock()
	if fake.ArtifactDirCreatedStub != nil {
		return fake.ArtifactDirCreatedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.artifactDirCreatedReturns.result1
}

func (fake *FakeInstance) ArtifactDirCreatedCallCount() int {
	fake.artifactDirCreatedMutex.RLock()
	defer fake.artifactDirCreatedMutex.RUnlock()
	return len(fake.artifactDirCreatedArgsForCall)
}

func (fake *FakeInstance) ArtifactDirCreatedReturns(result1 bool) {
	fake.ArtifactDirCreatedStub = nil
	fake.artifactDirCreatedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) ArtifactDirCreatedReturnsOnCall(i int, result1 bool) {
	fake.ArtifactDirCreatedStub = nil
	if fake.artifactDirCreatedReturnsOnCall == nil {
		fake.artifactDirCreatedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.artifactDirCreatedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) MarkArtifactDirCreated() {
	fake.markArtifactDirCreatedMutex.Lock()
	fake.markArtifactDirCreatedArgsForCall = append(fake.markArtifactDirCreatedArgsForCall, struct{}{})
	fake.recordInvocation("MarkArtifactDirCreated", []interface{}{})
	fake.markArtifactDirCreatedMutex.Unlock()
	if fake.MarkArtifactDirCreatedStub != nil {
		fake.MarkArtifactDirCreatedStub()
	}
}

func (fake *FakeInstance) MarkArtifactDirCreatedCallCount() int {
	fake.markArtifactDirCreatedMutex.RLock()
	defer fake.markArtifactDirCreatedMutex.RUnlock()
	return len(fake.markArtifactDirCreatedArgsForCall)
}

func (fake *FakeInstance) IsPostBackupUnlockable() bool {
	fake.isPostBackupUnlockableMutex.Lock()
	ret, specificReturn := fake.isPostBackupUnlockableReturnsOnCall[len(fake.isPostBackupUnlockableArgsForCall)]
	fake.isPostBackupUnlockableArgsForCall = append(fake.isPostBackupUnlockableArgsForCall, struct{}{})
	fake.recordInvocation("IsPostBackupUnlockable", []interface{}{})
	fake.isPostBackupUnlockableMutex.Unlock()
	if fake.IsPostBackupUnlockableStub != nil {
		return fake.IsPostBackupUnlockableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isPostBackupUnlockableReturns.result1
}

func (fake *FakeInstance) IsPostBackupUnlockableCallCount() int {
	fake.isPostBackupUnlockableMutex.RLock()
	defer fake.isPostBackupUnlockableMutex.RUnlock()
	return len(fake.isPostBackupUnlockableArgsForCall)
}

func (fake *FakeInstance) IsPostBackupUnlockableReturns(result1 bool) {
	fake.IsPostBackupUnlockableStub = nil
	fake.isPostBackupUnlockableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsPostBackupUnlockableReturnsOnCall(i int, result1 bool) {
	fake.IsPostBackupUnlockableStub = nil
	if fake.isPostBackupUnlockableReturnsOnCall == nil {
		fake.isPostBackupUnlockableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isPostBackupUnlockableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsPreBackupLockable() bool {
	fake.isPreBackupLockableMutex.Lock()
	ret, specificReturn := fake.isPreBackupLockableReturnsOnCall[len(fake.isPreBackupLockableArgsForCall)]
	fake.isPreBackupLockableArgsForCall = append(fake.isPreBackupLockableArgsForCall, struct{}{})
	fake.recordInvocation("IsPreBackupLockable", []interface{}{})
	fake.isPreBackupLockableMutex.Unlock()
	if fake.IsPreBackupLockableStub != nil {
		return fake.IsPreBackupLockableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isPreBackupLockableReturns.result1
}

func (fake *FakeInstance) IsPreBackupLockableCallCount() int {
	fake.isPreBackupLockableMutex.RLock()
	defer fake.isPreBackupLockableMutex.RUnlock()
	return len(fake.isPreBackupLockableArgsForCall)
}

func (fake *FakeInstance) IsPreBackupLockableReturns(result1 bool) {
	fake.IsPreBackupLockableStub = nil
	fake.isPreBackupLockableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsPreBackupLockableReturnsOnCall(i int, result1 bool) {
	fake.IsPreBackupLockableStub = nil
	if fake.isPreBackupLockableReturnsOnCall == nil {
		fake.isPreBackupLockableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isPreBackupLockableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsRestorable() bool {
	fake.isRestorableMutex.Lock()
	ret, specificReturn := fake.isRestorableReturnsOnCall[len(fake.isRestorableArgsForCall)]
	fake.isRestorableArgsForCall = append(fake.isRestorableArgsForCall, struct{}{})
	fake.recordInvocation("IsRestorable", []interface{}{})
	fake.isRestorableMutex.Unlock()
	if fake.IsRestorableStub != nil {
		return fake.IsRestorableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isRestorableReturns.result1
}

func (fake *FakeInstance) IsRestorableCallCount() int {
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	return len(fake.isRestorableArgsForCall)
}

func (fake *FakeInstance) IsRestorableReturns(result1 bool) {
	fake.IsRestorableStub = nil
	fake.isRestorableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsRestorableReturnsOnCall(i int, result1 bool) {
	fake.IsRestorableStub = nil
	if fake.isRestorableReturnsOnCall == nil {
		fake.isRestorableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isRestorableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) PreBackupLock() error {
	fake.preBackupLockMutex.Lock()
	ret, specificReturn := fake.preBackupLockReturnsOnCall[len(fake.preBackupLockArgsForCall)]
	fake.preBackupLockArgsForCall = append(fake.preBackupLockArgsForCall, struct{}{})
	fake.recordInvocation("PreBackupLock", []interface{}{})
	fake.preBackupLockMutex.Unlock()
	if fake.PreBackupLockStub != nil {
		return fake.PreBackupLockStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.preBackupLockReturns.result1
}

func (fake *FakeInstance) PreBackupLockCallCount() int {
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	return len(fake.preBackupLockArgsForCall)
}

func (fake *FakeInstance) PreBackupLockReturns(result1 error) {
	fake.PreBackupLockStub = nil
	fake.preBackupLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) PreBackupLockReturnsOnCall(i int, result1 error) {
	fake.PreBackupLockStub = nil
	if fake.preBackupLockReturnsOnCall == nil {
		fake.preBackupLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.preBackupLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Backup() error {
	fake.backupMutex.Lock()
	ret, specificReturn := fake.backupReturnsOnCall[len(fake.backupArgsForCall)]
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct{}{})
	fake.recordInvocation("Backup", []interface{}{})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.backupReturns.result1
}

func (fake *FakeInstance) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeInstance) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) BackupReturnsOnCall(i int, result1 error) {
	fake.BackupStub = nil
	if fake.backupReturnsOnCall == nil {
		fake.backupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.backupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) PostBackupUnlock() error {
	fake.postBackupUnlockMutex.Lock()
	ret, specificReturn := fake.postBackupUnlockReturnsOnCall[len(fake.postBackupUnlockArgsForCall)]
	fake.postBackupUnlockArgsForCall = append(fake.postBackupUnlockArgsForCall, struct{}{})
	fake.recordInvocation("PostBackupUnlock", []interface{}{})
	fake.postBackupUnlockMutex.Unlock()
	if fake.PostBackupUnlockStub != nil {
		return fake.PostBackupUnlockStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.postBackupUnlockReturns.result1
}

func (fake *FakeInstance) PostBackupUnlockCallCount() int {
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	return len(fake.postBackupUnlockArgsForCall)
}

func (fake *FakeInstance) PostBackupUnlockReturns(result1 error) {
	fake.PostBackupUnlockStub = nil
	fake.postBackupUnlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) PostBackupUnlockReturnsOnCall(i int, result1 error) {
	fake.PostBackupUnlockStub = nil
	if fake.postBackupUnlockReturnsOnCall == nil {
		fake.postBackupUnlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.postBackupUnlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Restore() error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct{}{})
	fake.recordInvocation("Restore", []interface{}{})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restoreReturns.result1
}

func (fake *FakeInstance) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeInstance) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) RestoreReturnsOnCall(i int, result1 error) {
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Cleanup() error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeInstance) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeInstance) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) CleanupPrevious() error {
	fake.cleanupPreviousMutex.Lock()
	ret, specificReturn := fake.cleanupPreviousReturnsOnCall[len(fake.cleanupPreviousArgsForCall)]
	fake.cleanupPreviousArgsForCall = append(fake.cleanupPreviousArgsForCall, struct{}{})
	fake.recordInvocation("CleanupPrevious", []interface{}{})
	fake.cleanupPreviousMutex.Unlock()
	if fake.CleanupPreviousStub != nil {
		return fake.CleanupPreviousStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupPreviousReturns.result1
}

func (fake *FakeInstance) CleanupPreviousCallCount() int {
	fake.cleanupPreviousMutex.RLock()
	defer fake.cleanupPreviousMutex.RUnlock()
	return len(fake.cleanupPreviousArgsForCall)
}

func (fake *FakeInstance) CleanupPreviousReturns(result1 error) {
	fake.CleanupPreviousStub = nil
	fake.cleanupPreviousReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) CleanupPreviousReturnsOnCall(i int, result1 error) {
	fake.CleanupPreviousStub = nil
	if fake.cleanupPreviousReturnsOnCall == nil {
		fake.cleanupPreviousReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupPreviousReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) ArtifactsToBackup() []orchestrator.BackupArtifact {
	fake.artifactsToBackupMutex.Lock()
	ret, specificReturn := fake.artifactsToBackupReturnsOnCall[len(fake.artifactsToBackupArgsForCall)]
	fake.artifactsToBackupArgsForCall = append(fake.artifactsToBackupArgsForCall, struct{}{})
	fake.recordInvocation("ArtifactsToBackup", []interface{}{})
	fake.artifactsToBackupMutex.Unlock()
	if fake.ArtifactsToBackupStub != nil {
		return fake.ArtifactsToBackupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.artifactsToBackupReturns.result1
}

func (fake *FakeInstance) ArtifactsToBackupCallCount() int {
	fake.artifactsToBackupMutex.RLock()
	defer fake.artifactsToBackupMutex.RUnlock()
	return len(fake.artifactsToBackupArgsForCall)
}

func (fake *FakeInstance) ArtifactsToBackupReturns(result1 []orchestrator.BackupArtifact) {
	fake.ArtifactsToBackupStub = nil
	fake.artifactsToBackupReturns = struct {
		result1 []orchestrator.BackupArtifact
	}{result1}
}

func (fake *FakeInstance) ArtifactsToBackupReturnsOnCall(i int, result1 []orchestrator.BackupArtifact) {
	fake.ArtifactsToBackupStub = nil
	if fake.artifactsToBackupReturnsOnCall == nil {
		fake.artifactsToBackupReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.BackupArtifact
		})
	}
	fake.artifactsToBackupReturnsOnCall[i] = struct {
		result1 []orchestrator.BackupArtifact
	}{result1}
}

func (fake *FakeInstance) ArtifactsToRestore() []orchestrator.BackupArtifact {
	fake.artifactsToRestoreMutex.Lock()
	ret, specificReturn := fake.artifactsToRestoreReturnsOnCall[len(fake.artifactsToRestoreArgsForCall)]
	fake.artifactsToRestoreArgsForCall = append(fake.artifactsToRestoreArgsForCall, struct{}{})
	fake.recordInvocation("ArtifactsToRestore", []interface{}{})
	fake.artifactsToRestoreMutex.Unlock()
	if fake.ArtifactsToRestoreStub != nil {
		return fake.ArtifactsToRestoreStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.artifactsToRestoreReturns.result1
}

func (fake *FakeInstance) ArtifactsToRestoreCallCount() int {
	fake.artifactsToRestoreMutex.RLock()
	defer fake.artifactsToRestoreMutex.RUnlock()
	return len(fake.artifactsToRestoreArgsForCall)
}

func (fake *FakeInstance) ArtifactsToRestoreReturns(result1 []orchestrator.BackupArtifact) {
	fake.ArtifactsToRestoreStub = nil
	fake.artifactsToRestoreReturns = struct {
		result1 []orchestrator.BackupArtifact
	}{result1}
}

func (fake *FakeInstance) ArtifactsToRestoreReturnsOnCall(i int, result1 []orchestrator.BackupArtifact) {
	fake.ArtifactsToRestoreStub = nil
	if fake.artifactsToRestoreReturnsOnCall == nil {
		fake.artifactsToRestoreReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.BackupArtifact
		})
	}
	fake.artifactsToRestoreReturnsOnCall[i] = struct {
		result1 []orchestrator.BackupArtifact
	}{result1}
}

func (fake *FakeInstance) CustomBackupArtifactNames() []string {
	fake.customBackupArtifactNamesMutex.Lock()
	ret, specificReturn := fake.customBackupArtifactNamesReturnsOnCall[len(fake.customBackupArtifactNamesArgsForCall)]
	fake.customBackupArtifactNamesArgsForCall = append(fake.customBackupArtifactNamesArgsForCall, struct{}{})
	fake.recordInvocation("CustomBackupArtifactNames", []interface{}{})
	fake.customBackupArtifactNamesMutex.Unlock()
	if fake.CustomBackupArtifactNamesStub != nil {
		return fake.CustomBackupArtifactNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.customBackupArtifactNamesReturns.result1
}

func (fake *FakeInstance) CustomBackupArtifactNamesCallCount() int {
	fake.customBackupArtifactNamesMutex.RLock()
	defer fake.customBackupArtifactNamesMutex.RUnlock()
	return len(fake.customBackupArtifactNamesArgsForCall)
}

func (fake *FakeInstance) CustomBackupArtifactNamesReturns(result1 []string) {
	fake.CustomBackupArtifactNamesStub = nil
	fake.customBackupArtifactNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) CustomBackupArtifactNamesReturnsOnCall(i int, result1 []string) {
	fake.CustomBackupArtifactNamesStub = nil
	if fake.customBackupArtifactNamesReturnsOnCall == nil {
		fake.customBackupArtifactNamesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.customBackupArtifactNamesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) CustomRestoreArtifactNames() []string {
	fake.customRestoreArtifactNamesMutex.Lock()
	ret, specificReturn := fake.customRestoreArtifactNamesReturnsOnCall[len(fake.customRestoreArtifactNamesArgsForCall)]
	fake.customRestoreArtifactNamesArgsForCall = append(fake.customRestoreArtifactNamesArgsForCall, struct{}{})
	fake.recordInvocation("CustomRestoreArtifactNames", []interface{}{})
	fake.customRestoreArtifactNamesMutex.Unlock()
	if fake.CustomRestoreArtifactNamesStub != nil {
		return fake.CustomRestoreArtifactNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.customRestoreArtifactNamesReturns.result1
}

func (fake *FakeInstance) CustomRestoreArtifactNamesCallCount() int {
	fake.customRestoreArtifactNamesMutex.RLock()
	defer fake.customRestoreArtifactNamesMutex.RUnlock()
	return len(fake.customRestoreArtifactNamesArgsForCall)
}

func (fake *FakeInstance) CustomRestoreArtifactNamesReturns(result1 []string) {
	fake.CustomRestoreArtifactNamesStub = nil
	fake.customRestoreArtifactNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) CustomRestoreArtifactNamesReturnsOnCall(i int, result1 []string) {
	fake.CustomRestoreArtifactNamesStub = nil
	if fake.customRestoreArtifactNamesReturnsOnCall == nil {
		fake.customRestoreArtifactNamesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.customRestoreArtifactNamesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	fake.artifactDirExistsMutex.RLock()
	defer fake.artifactDirExistsMutex.RUnlock()
	fake.artifactDirCreatedMutex.RLock()
	defer fake.artifactDirCreatedMutex.RUnlock()
	fake.markArtifactDirCreatedMutex.RLock()
	defer fake.markArtifactDirCreatedMutex.RUnlock()
	fake.isPostBackupUnlockableMutex.RLock()
	defer fake.isPostBackupUnlockableMutex.RUnlock()
	fake.isPreBackupLockableMutex.RLock()
	defer fake.isPreBackupLockableMutex.RUnlock()
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.cleanupPreviousMutex.RLock()
	defer fake.cleanupPreviousMutex.RUnlock()
	fake.artifactsToBackupMutex.RLock()
	defer fake.artifactsToBackupMutex.RUnlock()
	fake.artifactsToRestoreMutex.RLock()
	defer fake.artifactsToRestoreMutex.RUnlock()
	fake.customBackupArtifactNamesMutex.RLock()
	defer fake.customBackupArtifactNamesMutex.RUnlock()
	fake.customRestoreArtifactNamesMutex.RLock()
	defer fake.customRestoreArtifactNamesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstance) recordInvocation(key string, args []interface{}) {
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

var _ orchestrator.Instance = new(FakeInstance)
