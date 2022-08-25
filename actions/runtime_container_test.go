package actions

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedAction struct {
	mock.Mock
}

func (m *MockedAction) Identification() Identification {
	args := m.Called()

	return args.Get(0).(Identification)
}

func (m *MockedAction) Handle(a any) (any, error) {
	args := m.Called(a)

	return args.Get(0), args.Error(1)
}

func TestRuntimeContainer_Exists(t *testing.T) {
	kickActionIdentification := Identification("kick")

	cases := []struct {
		name   string
		init   func() Container
		action Identification
		expect func(result bool)
	}{
		{
			name: "Not exists with empty container",
			init: func() Container {
				return NewRuntimeContainer()
			},
			action: kickActionIdentification,
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "Not exists with filled container",
			init: func() Container {
				c := NewRuntimeContainer()
				mockedAction := &MockedAction{}
				mockedAction.On("Identification").
					Return(Identification("foo")).
					Once()
				c.Register(mockedAction)

				return c
			},
			action: kickActionIdentification,
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "Exists",
			init: func() Container {
				c := NewRuntimeContainer()
				mockedAction := &MockedAction{}
				mockedAction.On("Identification").
					Return(kickActionIdentification).
					Once()
				c.Register(mockedAction)

				return c
			},
			action: kickActionIdentification,
			expect: func(result bool) {
				assert.True(t, result)
			},
		},
		{
			name: "Not exists with filled container by init",
			init: func() Container {
				mockedAction := &MockedAction{}
				mockedAction.On("Identification").
					Return(Identification("foo")).
					Once()

				return NewRuntimeContainer(mockedAction)
			},
			action: kickActionIdentification,
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "Exists with filled container by init",
			init: func() Container {
				mockedAction := &MockedAction{}
				mockedAction.On("Identification").
					Return(kickActionIdentification).
					Once()

				return NewRuntimeContainer(mockedAction)
			},
			action: kickActionIdentification,
			expect: func(result bool) {
				assert.True(t, result)
			},
		},
	}

	for _, c := range cases {
		container := c.init()
		c.expect(container.Exists(c.action))
	}
}
