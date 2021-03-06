// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/noaa/events"
)

type FakeLogsNoaaRepository struct {
	GetContainerMetricsStub        func(string, []models.AppInstanceFields) ([]models.AppInstanceFields, error)
	getContainerMetricsMutex       sync.RWMutex
	getContainerMetricsArgsForCall []struct {
		arg1 string
		arg2 []models.AppInstanceFields
	}
	getContainerMetricsReturns struct {
		result1 []models.AppInstanceFields
		result2 error
	}
	RecentLogsForStub        func(appGuid string) ([]*events.LogMessage, error)
	recentLogsForMutex       sync.RWMutex
	recentLogsForArgsForCall []struct {
		appGuid string
	}
	recentLogsForReturns struct {
		result1 []*events.LogMessage
		result2 error
	}
	TailNoaaLogsForStub        func(appGuid string, onConnect func(), onMessage func(*events.LogMessage)) error
	tailNoaaLogsForMutex       sync.RWMutex
	tailNoaaLogsForArgsForCall []struct {
		appGuid   string
		onConnect func()
		onMessage func(*events.LogMessage)
	}
	tailNoaaLogsForReturns struct {
		result1 error
	}
	NewConsumerStub        func(api.NoaaConsumer)
	newConsumerMutex       sync.RWMutex
	newConsumerArgsForCall []struct {
		arg1 api.NoaaConsumer
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
}

func (fake *FakeLogsNoaaRepository) GetContainerMetrics(arg1 string, arg2 []models.AppInstanceFields) ([]models.AppInstanceFields, error) {
	fake.getContainerMetricsMutex.Lock()
	defer fake.getContainerMetricsMutex.Unlock()
	fake.getContainerMetricsArgsForCall = append(fake.getContainerMetricsArgsForCall, struct {
		arg1 string
		arg2 []models.AppInstanceFields
	}{arg1, arg2})
	if fake.GetContainerMetricsStub != nil {
		return fake.GetContainerMetricsStub(arg1, arg2)
	} else {
		return fake.getContainerMetricsReturns.result1, fake.getContainerMetricsReturns.result2
	}
}

func (fake *FakeLogsNoaaRepository) GetContainerMetricsCallCount() int {
	fake.getContainerMetricsMutex.RLock()
	defer fake.getContainerMetricsMutex.RUnlock()
	return len(fake.getContainerMetricsArgsForCall)
}

func (fake *FakeLogsNoaaRepository) GetContainerMetricsArgsForCall(i int) (string, []models.AppInstanceFields) {
	fake.getContainerMetricsMutex.RLock()
	defer fake.getContainerMetricsMutex.RUnlock()
	return fake.getContainerMetricsArgsForCall[i].arg1, fake.getContainerMetricsArgsForCall[i].arg2
}

func (fake *FakeLogsNoaaRepository) GetContainerMetricsReturns(result1 []models.AppInstanceFields, result2 error) {
	fake.GetContainerMetricsStub = nil
	fake.getContainerMetricsReturns = struct {
		result1 []models.AppInstanceFields
		result2 error
	}{result1, result2}
}

func (fake *FakeLogsNoaaRepository) RecentLogsFor(appGuid string) ([]*events.LogMessage, error) {
	fake.recentLogsForMutex.Lock()
	defer fake.recentLogsForMutex.Unlock()
	fake.recentLogsForArgsForCall = append(fake.recentLogsForArgsForCall, struct {
		appGuid string
	}{appGuid})
	if fake.RecentLogsForStub != nil {
		return fake.RecentLogsForStub(appGuid)
	} else {
		return fake.recentLogsForReturns.result1, fake.recentLogsForReturns.result2
	}
}

func (fake *FakeLogsNoaaRepository) RecentLogsForCallCount() int {
	fake.recentLogsForMutex.RLock()
	defer fake.recentLogsForMutex.RUnlock()
	return len(fake.recentLogsForArgsForCall)
}

func (fake *FakeLogsNoaaRepository) RecentLogsForArgsForCall(i int) string {
	fake.recentLogsForMutex.RLock()
	defer fake.recentLogsForMutex.RUnlock()
	return fake.recentLogsForArgsForCall[i].appGuid
}

func (fake *FakeLogsNoaaRepository) RecentLogsForReturns(result1 []*events.LogMessage, result2 error) {
	fake.RecentLogsForStub = nil
	fake.recentLogsForReturns = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeLogsNoaaRepository) TailNoaaLogsFor(appGuid string, onConnect func(), onMessage func(*events.LogMessage)) error {
	fake.tailNoaaLogsForMutex.Lock()
	defer fake.tailNoaaLogsForMutex.Unlock()
	fake.tailNoaaLogsForArgsForCall = append(fake.tailNoaaLogsForArgsForCall, struct {
		appGuid   string
		onConnect func()
		onMessage func(*events.LogMessage)
	}{appGuid, onConnect, onMessage})
	if fake.TailNoaaLogsForStub != nil {
		return fake.TailNoaaLogsForStub(appGuid, onConnect, onMessage)
	} else {
		return fake.tailNoaaLogsForReturns.result1
	}
}

func (fake *FakeLogsNoaaRepository) TailNoaaLogsForCallCount() int {
	fake.tailNoaaLogsForMutex.RLock()
	defer fake.tailNoaaLogsForMutex.RUnlock()
	return len(fake.tailNoaaLogsForArgsForCall)
}

func (fake *FakeLogsNoaaRepository) TailNoaaLogsForArgsForCall(i int) (string, func(), func(*events.LogMessage)) {
	fake.tailNoaaLogsForMutex.RLock()
	defer fake.tailNoaaLogsForMutex.RUnlock()
	return fake.tailNoaaLogsForArgsForCall[i].appGuid, fake.tailNoaaLogsForArgsForCall[i].onConnect, fake.tailNoaaLogsForArgsForCall[i].onMessage
}

func (fake *FakeLogsNoaaRepository) TailNoaaLogsForReturns(result1 error) {
	fake.TailNoaaLogsForStub = nil
	fake.tailNoaaLogsForReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLogsNoaaRepository) NewConsumer(arg1 api.NoaaConsumer) {
	fake.newConsumerMutex.Lock()
	defer fake.newConsumerMutex.Unlock()
	fake.newConsumerArgsForCall = append(fake.newConsumerArgsForCall, struct {
		arg1 api.NoaaConsumer
	}{arg1})
	if fake.NewConsumerStub != nil {
		fake.NewConsumerStub(arg1)
	}
}

func (fake *FakeLogsNoaaRepository) NewConsumerCallCount() int {
	fake.newConsumerMutex.RLock()
	defer fake.newConsumerMutex.RUnlock()
	return len(fake.newConsumerArgsForCall)
}

func (fake *FakeLogsNoaaRepository) NewConsumerArgsForCall(i int) api.NoaaConsumer {
	fake.newConsumerMutex.RLock()
	defer fake.newConsumerMutex.RUnlock()
	return fake.newConsumerArgsForCall[i].arg1
}

func (fake *FakeLogsNoaaRepository) Close() {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeLogsNoaaRepository) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

var _ api.LogsNoaaRepository = new(FakeLogsNoaaRepository)
