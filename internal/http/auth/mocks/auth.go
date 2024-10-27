// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	models "kudago/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// CheckCredentials mocks base method.
func (m *MockAuthService) CheckCredentials(ctx context.Context, creds models.Credentials) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCredentials", ctx, creds)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCredentials indicates an expected call of CheckCredentials.
func (mr *MockAuthServiceMockRecorder) CheckCredentials(ctx, creds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCredentials", reflect.TypeOf((*MockAuthService)(nil).CheckCredentials), ctx, creds)
}

// CheckSession mocks base method.
func (m *MockAuthService) CheckSession(ctx context.Context, cookie string) (models.Session, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, cookie)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockAuthServiceMockRecorder) CheckSession(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockAuthService)(nil).CheckSession), ctx, cookie)
}

// CreateSession mocks base method.
func (m *MockAuthService) CreateSession(ctx context.Context, ID int) (models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, ID)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthServiceMockRecorder) CreateSession(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthService)(nil).CreateSession), ctx, ID)
}

// DeleteSession mocks base method.
func (m *MockAuthService) DeleteSession(ctx context.Context, token string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteSession", ctx, token)
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockAuthServiceMockRecorder) DeleteSession(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockAuthService)(nil).DeleteSession), ctx, token)
}

// GetUserByID mocks base method.
func (m *MockAuthService) GetUserByID(ctx context.Context, ID int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, ID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthServiceMockRecorder) GetUserByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuthService)(nil).GetUserByID), ctx, ID)
}

// Register mocks base method.
func (m *MockAuthService) Register(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthService)(nil).Register), ctx, user)
}