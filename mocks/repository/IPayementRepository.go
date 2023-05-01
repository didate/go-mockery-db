// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	models "mockerydb/models"

	mock "github.com/stretchr/testify/mock"
)

// IPayementRepository is an autogenerated mock type for the IPayementRepository type
type IPayementRepository struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: payment
func (_m *IPayementRepository) CreatePayment(payment models.Payment) (int64, error) {
	ret := _m.Called(payment)

	var r0 int64
	if rf, ok := ret.Get(0).(func(models.Payment) int64); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Payment) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *IPayementRepository) DeletePayment(id string) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectPaymentWithId provides a mock function with given fields: id
func (_m *IPayementRepository) SelectPaymentWithId(id string) (models.Payment, error) {
	ret := _m.Called(id)

	var r0 models.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Payment, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Payment); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Payment)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayment provides a mock function with given fields: id, payment
func (_m *IPayementRepository) UpdatePayment(id string, payment models.Payment) (models.Payment, error) {
	ret := _m.Called(id, payment)

	var r0 models.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Payment) (models.Payment, error)); ok {
		return rf(id, payment)
	}
	if rf, ok := ret.Get(0).(func(string, models.Payment) models.Payment); ok {
		r0 = rf(id, payment)
	} else {
		r0 = ret.Get(0).(models.Payment)
	}

	if rf, ok := ret.Get(1).(func(string, models.Payment) error); ok {
		r1 = rf(id, payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIPayementRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPayementRepository creates a new instance of IPayementRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPayementRepository(t mockConstructorTestingTNewIPayementRepository) *IPayementRepository {
	mock := &IPayementRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
