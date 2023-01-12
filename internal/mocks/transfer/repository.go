// Code generated by mockery v2.15.0. DO NOT EDIT.

package transfer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	internaltransfer "github.com/willkoerich/rock-challenge/internal/transfer"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]internaltransfer.Transfer, error) {
	ret := _m.Called(ctx)

	var r0 []internaltransfer.Transfer
	if rf, ok := ret.Get(0).(func(context.Context) []internaltransfer.Transfer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internaltransfer.Transfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAccountOriginID provides a mock function with given fields: ctx, accountOriginID
func (_m *Repository) GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]internaltransfer.Transfer, error) {
	ret := _m.Called(ctx, accountOriginID)

	var r0 []internaltransfer.Transfer
	if rf, ok := ret.Get(0).(func(context.Context, int) []internaltransfer.Transfer); ok {
		r0 = rf(ctx, accountOriginID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internaltransfer.Transfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, accountOriginID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int) (internaltransfer.Transfer, error) {
	ret := _m.Called(ctx, id)

	var r0 internaltransfer.Transfer
	if rf, ok := ret.Get(0).(func(context.Context, int) internaltransfer.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(internaltransfer.Transfer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, transfer
func (_m *Repository) Save(ctx context.Context, transfer internaltransfer.Transfer) (internaltransfer.Transfer, error) {
	ret := _m.Called(ctx, transfer)

	var r0 internaltransfer.Transfer
	if rf, ok := ret.Get(0).(func(context.Context, internaltransfer.Transfer) internaltransfer.Transfer); ok {
		r0 = rf(ctx, transfer)
	} else {
		r0 = ret.Get(0).(internaltransfer.Transfer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, internaltransfer.Transfer) error); ok {
		r1 = rf(ctx, transfer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
