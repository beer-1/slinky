// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider[K comparable, V interface{}] struct {
	mock.Mock
}

// GetData provides a mock function with given fields: _a0
func (_m *Provider[K, V]) GetData(_a0 context.Context) (map[K]V, error) {
	ret := _m.Called(_a0)

	var r0 map[K]V
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[K]V, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[K]V); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[K]V)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastUpdate provides a mock function with given fields:
func (_m *Provider[K, V]) LastUpdate() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Provider[K, V]) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *Provider[K, V]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvider[K comparable, V interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Provider[K, V] {
	mock := &Provider[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}