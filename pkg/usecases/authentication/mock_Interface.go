// Code generated by mockery v2.13.1. DO NOT EDIT.

package authentication

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: ctx, in
func (_m *MockInterface) Do(ctx context.Context, in Input) (*Output, error) {
	ret := _m.Called(ctx, in)

	var r0 *Output
	if rf, ok := ret.Get(0).(func(context.Context, Input) *Output); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Input) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type MockInterface_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//  - ctx context.Context
//  - in Input
func (_e *MockInterface_Expecter) Do(ctx interface{}, in interface{}) *MockInterface_Do_Call {
	return &MockInterface_Do_Call{Call: _e.mock.On("Do", ctx, in)}
}

func (_c *MockInterface_Do_Call) Run(run func(ctx context.Context, in Input)) *MockInterface_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Input))
	})
	return _c
}

func (_c *MockInterface_Do_Call) Return(_a0 *Output, _a1 error) *MockInterface_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterface(t mockConstructorTestingTNewMockInterface) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
