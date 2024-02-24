// Code generated by mockery v2.42.0. DO NOT EDIT.

package application

import (
	context "context"
	domain "quizazz/quiz-service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockApp is an autogenerated mock type for the App type
type MockApp struct {
	mock.Mock
}

// AuthorizeQuiz provides a mock function with given fields: ctx, authorize
func (_m *MockApp) AuthorizeQuiz(ctx context.Context, authorize AuthorizeQuiz) error {
	ret := _m.Called(ctx, authorize)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizeQuiz")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AuthorizeQuiz) error); ok {
		r0 = rf(ctx, authorize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateQuiz provides a mock function with given fields: ctx, register
func (_m *MockApp) CreateQuiz(ctx context.Context, register CreateQuiz) error {
	ret := _m.Called(ctx, register)

	if len(ret) == 0 {
		panic("no return value specified for CreateQuiz")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateQuiz) error); ok {
		r0 = rf(ctx, register)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableQuiz provides a mock function with given fields: ctx, disable
func (_m *MockApp) DisableQuiz(ctx context.Context, disable DisableQuiz) error {
	ret := _m.Called(ctx, disable)

	if len(ret) == 0 {
		panic("no return value specified for DisableQuiz")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, DisableQuiz) error); ok {
		r0 = rf(ctx, disable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableQuiz provides a mock function with given fields: ctx, enable
func (_m *MockApp) EnableQuiz(ctx context.Context, enable EnableQuiz) error {
	ret := _m.Called(ctx, enable)

	if len(ret) == 0 {
		panic("no return value specified for EnableQuiz")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, EnableQuiz) error); ok {
		r0 = rf(ctx, enable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetQuiz provides a mock function with given fields: ctx, get
func (_m *MockApp) GetQuiz(ctx context.Context, get GetQuiz) (*domain.Quiz, error) {
	ret := _m.Called(ctx, get)

	if len(ret) == 0 {
		panic("no return value specified for GetQuiz")
	}

	var r0 *domain.Quiz
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, GetQuiz) (*domain.Quiz, error)); ok {
		return rf(ctx, get)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GetQuiz) *domain.Quiz); ok {
		r0 = rf(ctx, get)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Quiz)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GetQuiz) error); ok {
		r1 = rf(ctx, get)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockApp creates a new instance of MockApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockApp(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockApp {
	mock := &MockApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
