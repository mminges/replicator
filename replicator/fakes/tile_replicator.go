// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mminges/replicator/replicator"
)

type TileReplicator struct {
	ReplicateStub        func(replicator.ApplicationConfig) error
	replicateMutex       sync.RWMutex
	replicateArgsForCall []struct {
		arg1 replicator.ApplicationConfig
	}
	replicateReturns struct {
		result1 error
	}
	replicateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TileReplicator) Replicate(arg1 replicator.ApplicationConfig) error {
	fake.replicateMutex.Lock()
	ret, specificReturn := fake.replicateReturnsOnCall[len(fake.replicateArgsForCall)]
	fake.replicateArgsForCall = append(fake.replicateArgsForCall, struct {
		arg1 replicator.ApplicationConfig
	}{arg1})
	fake.recordInvocation("Replicate", []interface{}{arg1})
	fake.replicateMutex.Unlock()
	if fake.ReplicateStub != nil {
		return fake.ReplicateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.replicateReturns.result1
}

func (fake *TileReplicator) ReplicateCallCount() int {
	fake.replicateMutex.RLock()
	defer fake.replicateMutex.RUnlock()
	return len(fake.replicateArgsForCall)
}

func (fake *TileReplicator) ReplicateArgsForCall(i int) replicator.ApplicationConfig {
	fake.replicateMutex.RLock()
	defer fake.replicateMutex.RUnlock()
	return fake.replicateArgsForCall[i].arg1
}

func (fake *TileReplicator) ReplicateReturns(result1 error) {
	fake.ReplicateStub = nil
	fake.replicateReturns = struct {
		result1 error
	}{result1}
}

func (fake *TileReplicator) ReplicateReturnsOnCall(i int, result1 error) {
	fake.ReplicateStub = nil
	if fake.replicateReturnsOnCall == nil {
		fake.replicateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.replicateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TileReplicator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.replicateMutex.RLock()
	defer fake.replicateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TileReplicator) recordInvocation(key string, args []interface{}) {
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
