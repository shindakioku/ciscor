package runner

import (
	"context"
	"github.com/shindakioku/ciscor/actions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

type MockedAction struct {
	mock.Mock
	waitGroup  *sync.WaitGroup
	timeCalls  int
	passedArgs any
}

func (m *MockedAction) Identification() actions.Identification {
	args := m.Called()

	return args.Get(0).(actions.Identification)
}

func (m *MockedAction) Handle(a any) (any, error) {
	args := m.Called(a)
	m.passedArgs = a
	m.timeCalls += 1
	m.waitGroup.Done()

	return args.Get(0), args.Error(1)
}

func TestSimpleASyncRunner_Run(t *testing.T) {
	asyncRunner := NewSimpleASyncRunner(10)
	ctx := context.Background()
	go asyncRunner.Run(ctx)

	var waitGroup sync.WaitGroup
	mockedAction := MockedAction{waitGroup: &waitGroup}
	mockedAction.On("Handle", mock.Anything).
		Return(nil, nil).
		Once()
	asyncRunner.Add(AsyncJob{
		Action: &mockedAction,
		Args:   "hello world!",
	})
	waitGroup.Add(1)
	waitGroup.Wait()
	ctx.Done()

	assert.Equal(t, "hello world!", mockedAction.passedArgs)
	assert.Equal(t, 1, mockedAction.timeCalls)
}
