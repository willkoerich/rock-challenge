// Code generated by mockery v2.15.0. DO NOT EDIT.

package database

import mock "github.com/stretchr/testify/mock"

// Results is an autogenerated mock type for the Results type
type Results struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Results) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Results) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *Results) Scan(dest ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dest...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewResults interface {
	mock.TestingT
	Cleanup(func())
}

// NewResults creates a new instance of Results. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResults(t mockConstructorTestingTNewResults) *Results {
	mock := &Results{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}