// Code generated by mockery v2.13.1. DO NOT EDIT.

package logger

import mock "github.com/stretchr/testify/mock"

// mockGoLogger is an autogenerated mock type for the goLogger type
type mockGoLogger struct {
	mock.Mock
}

type mockGoLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *mockGoLogger) EXPECT() *mockGoLogger_Expecter {
	return &mockGoLogger_Expecter{mock: &_m.Mock}
}

// Printf provides a mock function with given fields: format, v
func (_m *mockGoLogger) Printf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// mockGoLogger_Printf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Printf'
type mockGoLogger_Printf_Call struct {
	*mock.Call
}

// Printf is a helper method to define mock.On call
//  - format string
//  - v ...interface{}
func (_e *mockGoLogger_Expecter) Printf(format interface{}, v ...interface{}) *mockGoLogger_Printf_Call {
	return &mockGoLogger_Printf_Call{Call: _e.mock.On("Printf",
		append([]interface{}{format}, v...)...)}
}

func (_c *mockGoLogger_Printf_Call) Run(run func(format string, v ...interface{})) *mockGoLogger_Printf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *mockGoLogger_Printf_Call) Return() *mockGoLogger_Printf_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockGoLogger interface {
	mock.TestingT
	Cleanup(func())
}

// newMockGoLogger creates a new instance of mockGoLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockGoLogger(t mockConstructorTestingTnewMockGoLogger) *mockGoLogger {
	mock := &mockGoLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
