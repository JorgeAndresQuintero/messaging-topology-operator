// Code generated by counterfeiter. DO NOT EDIT.
package rabbitmqclientfakes

import (
	"sync"

	"github.com/hashicorp/vault/api"
	"github.com/rabbitmq/messaging-topology-operator/rabbitmqclient"
)

type FakeSecretReader struct {
	ReadSecretStub        func(string) (*api.Secret, error)
	readSecretMutex       sync.RWMutex
	readSecretArgsForCall []struct {
		arg1 string
	}
	readSecretReturns struct {
		result1 *api.Secret
		result2 error
	}
	readSecretReturnsOnCall map[int]struct {
		result1 *api.Secret
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecretReader) ReadSecret(arg1 string) (*api.Secret, error) {
	fake.readSecretMutex.Lock()
	ret, specificReturn := fake.readSecretReturnsOnCall[len(fake.readSecretArgsForCall)]
	fake.readSecretArgsForCall = append(fake.readSecretArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReadSecretStub
	fakeReturns := fake.readSecretReturns
	fake.recordInvocation("ReadSecret", []interface{}{arg1})
	fake.readSecretMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretReader) ReadSecretCallCount() int {
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	return len(fake.readSecretArgsForCall)
}

func (fake *FakeSecretReader) ReadSecretCalls(stub func(string) (*api.Secret, error)) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = stub
}

func (fake *FakeSecretReader) ReadSecretArgsForCall(i int) string {
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	argsForCall := fake.readSecretArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecretReader) ReadSecretReturns(result1 *api.Secret, result2 error) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = nil
	fake.readSecretReturns = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretReader) ReadSecretReturnsOnCall(i int, result1 *api.Secret, result2 error) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = nil
	if fake.readSecretReturnsOnCall == nil {
		fake.readSecretReturnsOnCall = make(map[int]struct {
			result1 *api.Secret
			result2 error
		})
	}
	fake.readSecretReturnsOnCall[i] = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecretReader) recordInvocation(key string, args []interface{}) {
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

var _ rabbitmqclient.SecretReader = new(FakeSecretReader)
