// Code generated by mockery v2.15.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	internaldomain "github.com/willkoerich/rock-challenge/internal/domain"
)

// LoginController is an autogenerated mock type for the LoginController type
type LoginController struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, credential
func (_m *LoginController) Authenticate(ctx context.Context, credential internaldomain.AuthenticationRequest) (internaldomain.AuthenticationResponse, error) {
	ret := _m.Called(ctx, credential)

	var r0 internaldomain.AuthenticationResponse
	if rf, ok := ret.Get(0).(func(context.Context, internaldomain.AuthenticationRequest) internaldomain.AuthenticationResponse); ok {
		r0 = rf(ctx, credential)
	} else {
		r0 = ret.Get(0).(internaldomain.AuthenticationResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, internaldomain.AuthenticationRequest) error); ok {
		r1 = rf(ctx, credential)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLoginController interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginController creates a new instance of LoginController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginController(t mockConstructorTestingTNewLoginController) *LoginController {
	mock := &LoginController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
