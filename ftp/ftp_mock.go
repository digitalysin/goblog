// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package ftp

import (
	"io"

	mock "github.com/stretchr/testify/mock"
)

// NewMockFtp creates a new instance of MockFtp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFtp(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFtp {
	mock := &MockFtp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockFtp is an autogenerated mock type for the Ftp type
type MockFtp struct {
	mock.Mock
}

type MockFtp_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFtp) EXPECT() *MockFtp_Expecter {
	return &MockFtp_Expecter{mock: &_m.Mock}
}

// Close provides a mock function for the type MockFtp
func (_mock *MockFtp) Close() error {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func() error); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockFtp_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockFtp_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockFtp_Expecter) Close() *MockFtp_Close_Call {
	return &MockFtp_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockFtp_Close_Call) Run(run func()) *MockFtp_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFtp_Close_Call) Return(err error) *MockFtp_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFtp_Close_Call) RunAndReturn(run func() error) *MockFtp_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockFtp
func (_mock *MockFtp) Get(path string, dst io.Writer) error {
	ret := _mock.Called(path, dst)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(string, io.Writer) error); ok {
		r0 = returnFunc(path, dst)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockFtp_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockFtp_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path
//   - dst
func (_e *MockFtp_Expecter) Get(path interface{}, dst interface{}) *MockFtp_Get_Call {
	return &MockFtp_Get_Call{Call: _e.mock.On("Get", path, dst)}
}

func (_c *MockFtp_Get_Call) Run(run func(path string, dst io.Writer)) *MockFtp_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(io.Writer))
	})
	return _c
}

func (_c *MockFtp_Get_Call) Return(err error) *MockFtp_Get_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFtp_Get_Call) RunAndReturn(run func(path string, dst io.Writer) error) *MockFtp_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function for the type MockFtp
func (_mock *MockFtp) List(path string) ([]Entry, error) {
	ret := _mock.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []Entry
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string) ([]Entry, error)); ok {
		return returnFunc(path)
	}
	if returnFunc, ok := ret.Get(0).(func(string) []Entry); ok {
		r0 = returnFunc(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Entry)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string) error); ok {
		r1 = returnFunc(path)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockFtp_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockFtp_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - path
func (_e *MockFtp_Expecter) List(path interface{}) *MockFtp_List_Call {
	return &MockFtp_List_Call{Call: _e.mock.On("List", path)}
}

func (_c *MockFtp_List_Call) Run(run func(path string)) *MockFtp_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFtp_List_Call) Return(entrys []Entry, err error) *MockFtp_List_Call {
	_c.Call.Return(entrys, err)
	return _c
}

func (_c *MockFtp_List_Call) RunAndReturn(run func(path string) ([]Entry, error)) *MockFtp_List_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function for the type MockFtp
func (_mock *MockFtp) Put(path string, src io.Reader) error {
	ret := _mock.Called(path, src)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = returnFunc(path, src)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockFtp_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockFtp_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - path
//   - src
func (_e *MockFtp_Expecter) Put(path interface{}, src interface{}) *MockFtp_Put_Call {
	return &MockFtp_Put_Call{Call: _e.mock.On("Put", path, src)}
}

func (_c *MockFtp_Put_Call) Run(run func(path string, src io.Reader)) *MockFtp_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(io.Reader))
	})
	return _c
}

func (_c *MockFtp_Put_Call) Return(err error) *MockFtp_Put_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFtp_Put_Call) RunAndReturn(run func(path string, src io.Reader) error) *MockFtp_Put_Call {
	_c.Call.Return(run)
	return _c
}
