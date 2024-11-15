// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "kudago/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserDB is a mock of UserDB interface.
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB.
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates a new mock instance.
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserDB) AddUser(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserDBMockRecorder) AddUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserDB)(nil).AddUser), ctx, user)
}

// CheckCredentials mocks base method.
func (m *MockUserDB) CheckCredentials(ctx context.Context, username, password string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCredentials", ctx, username, password)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCredentials indicates an expected call of CheckCredentials.
func (mr *MockUserDBMockRecorder) CheckCredentials(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCredentials", reflect.TypeOf((*MockUserDB)(nil).CheckCredentials), ctx, username, password)
}

// CheckEmail mocks base method.
func (m *MockUserDB) CheckEmail(ctx context.Context, email string, ID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmail", ctx, email, ID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmail indicates an expected call of CheckEmail.
func (mr *MockUserDBMockRecorder) CheckEmail(ctx, email, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmail", reflect.TypeOf((*MockUserDB)(nil).CheckEmail), ctx, email, ID)
}

// CheckUsername mocks base method.
func (m *MockUserDB) CheckUsername(ctx context.Context, username string, ID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsername", ctx, username, ID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsername indicates an expected call of CheckUsername.
func (mr *MockUserDBMockRecorder) CheckUsername(ctx, username, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsername", reflect.TypeOf((*MockUserDB)(nil).CheckUsername), ctx, username, ID)
}

// GetUserByID mocks base method.
func (m *MockUserDB) GetUserByID(ctx context.Context, ID int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, ID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserDBMockRecorder) GetUserByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserDB)(nil).GetUserByID), ctx, ID)
}

// Subscribe mocks base method.
func (m *MockUserDB) Subscribe(ctx context.Context, subscription models.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockUserDBMockRecorder) Subscribe(ctx, subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockUserDB)(nil).Subscribe), ctx, subscription)
}

// Unsubscribe mocks base method.
func (m *MockUserDB) Unsubscribe(ctx context.Context, subscription models.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx, subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockUserDBMockRecorder) Unsubscribe(ctx, subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockUserDB)(nil).Unsubscribe), ctx, subscription)
}

// UpdateUser mocks base method.
func (m *MockUserDB) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserDBMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserDB)(nil).UpdateUser), ctx, user)
}

// UserExists mocks base method.
func (m *MockUserDB) UserExists(ctx context.Context, username, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserExists", ctx, username, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserExists indicates an expected call of UserExists.
func (mr *MockUserDBMockRecorder) UserExists(ctx, username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserExists", reflect.TypeOf((*MockUserDB)(nil).UserExists), ctx, username, email)
}

// MockSessionDB is a mock of SessionDB interface.
type MockSessionDB struct {
	ctrl     *gomock.Controller
	recorder *MockSessionDBMockRecorder
}

// MockSessionDBMockRecorder is the mock recorder for MockSessionDB.
type MockSessionDBMockRecorder struct {
	mock *MockSessionDB
}

// NewMockSessionDB creates a new mock instance.
func NewMockSessionDB(ctrl *gomock.Controller) *MockSessionDB {
	mock := &MockSessionDB{ctrl: ctrl}
	mock.recorder = &MockSessionDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionDB) EXPECT() *MockSessionDBMockRecorder {
	return m.recorder
}

// CheckSession mocks base method.
func (m *MockSessionDB) CheckSession(ctx context.Context, cookie string) (models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, cookie)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockSessionDBMockRecorder) CheckSession(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockSessionDB)(nil).CheckSession), ctx, cookie)
}

// CreateSession mocks base method.
func (m *MockSessionDB) CreateSession(ctx context.Context, ID int) (models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, ID)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionDBMockRecorder) CreateSession(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionDB)(nil).CreateSession), ctx, ID)
}

// DeleteSession mocks base method.
func (m *MockSessionDB) DeleteSession(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionDBMockRecorder) DeleteSession(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionDB)(nil).DeleteSession), ctx, token)
}

// MockImageDB is a mock of ImageDB interface.
type MockImageDB struct {
	ctrl     *gomock.Controller
	recorder *MockImageDBMockRecorder
}

// MockImageDBMockRecorder is the mock recorder for MockImageDB.
type MockImageDBMockRecorder struct {
	mock *MockImageDB
}

// NewMockImageDB creates a new mock instance.
func NewMockImageDB(ctrl *gomock.Controller) *MockImageDB {
	mock := &MockImageDB{ctrl: ctrl}
	mock.recorder = &MockImageDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageDB) EXPECT() *MockImageDBMockRecorder {
	return m.recorder
}

// DeleteImage mocks base method.
func (m *MockImageDB) DeleteImage(ctx context.Context, imagePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", ctx, imagePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockImageDBMockRecorder) DeleteImage(ctx, imagePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockImageDB)(nil).DeleteImage), ctx, imagePath)
}

// SaveImage mocks base method.
func (m *MockImageDB) SaveImage(ctx context.Context, media models.MediaFile) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveImage", ctx, media)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveImage indicates an expected call of SaveImage.
func (mr *MockImageDBMockRecorder) SaveImage(ctx, media interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveImage", reflect.TypeOf((*MockImageDB)(nil).SaveImage), ctx, media)
}
