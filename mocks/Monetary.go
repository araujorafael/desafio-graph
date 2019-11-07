// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Monetary is an autogenerated mock type for the Monetary type
type Monetary struct {
	mock.Mock
}

// Convert provides a mock function with given fields: value, currency
func (_m *Monetary) Convert(value float64, currency string) float64 {
	ret := _m.Called(value, currency)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64, string) float64); ok {
		r0 = rf(value, currency)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}