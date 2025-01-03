// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/usecase/user.go

// Package mock is a generated GoMock package.
package mock

import (
	models "go-fiber-v2/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecaseService is a mock of UserUsecaseService interface.
type MockUserUsecaseService struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseServiceMockRecorder
}

// MockUserUsecaseServiceMockRecorder is the mock recorder for MockUserUsecaseService.
type MockUserUsecaseServiceMockRecorder struct {
	mock *MockUserUsecaseService
}

// NewMockUserUsecaseService creates a new mock instance.
func NewMockUserUsecaseService(ctrl *gomock.Controller) *MockUserUsecaseService {
	mock := &MockUserUsecaseService{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecaseService) EXPECT() *MockUserUsecaseServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUsecaseService) CreateUser(up *models.SignUpRequest) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", up)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUsecaseServiceMockRecorder) CreateUser(up interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUsecaseService)(nil).CreateUser), up)
}
