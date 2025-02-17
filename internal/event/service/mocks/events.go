// Code generated by MockGen. DO NOT EDIT.
// Source: ./event.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "kudago/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventDB is a mock of EventDB interface.
type MockEventDB struct {
	ctrl     *gomock.Controller
	recorder *MockEventDBMockRecorder
}

// MockEventDBMockRecorder is the mock recorder for MockEventDB.
type MockEventDBMockRecorder struct {
	mock *MockEventDB
}

// NewMockEventDB creates a new mock instance.
func NewMockEventDB(ctrl *gomock.Controller) *MockEventDB {
	mock := &MockEventDB{ctrl: ctrl}
	mock.recorder = &MockEventDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventDB) EXPECT() *MockEventDBMockRecorder {
	return m.recorder
}

// AddEventToFavorites mocks base method.
func (m *MockEventDB) AddEventToFavorites(ctx context.Context, newFavorite models.FavoriteEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventToFavorites", ctx, newFavorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventToFavorites indicates an expected call of AddEventToFavorites.
func (mr *MockEventDBMockRecorder) AddEventToFavorites(ctx, newFavorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventToFavorites", reflect.TypeOf((*MockEventDB)(nil).AddEventToFavorites), ctx, newFavorite)
}

// CreateEvent mocks base method.
func (m *MockEventDB) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, event)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockEventDBMockRecorder) CreateEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEventDB)(nil).CreateEvent), ctx, event)
}

// DeleteEvent mocks base method.
func (m *MockEventDB) DeleteEvent(ctx context.Context, ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockEventDBMockRecorder) DeleteEvent(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockEventDB)(nil).DeleteEvent), ctx, ID)
}

// DeleteEventFromFavorites mocks base method.
func (m *MockEventDB) DeleteEventFromFavorites(ctx context.Context, favorite models.FavoriteEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventFromFavorites", ctx, favorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventFromFavorites indicates an expected call of DeleteEventFromFavorites.
func (mr *MockEventDBMockRecorder) DeleteEventFromFavorites(ctx, favorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventFromFavorites", reflect.TypeOf((*MockEventDB)(nil).DeleteEventFromFavorites), ctx, favorite)
}

// GetCategories mocks base method.
func (m *MockEventDB) GetCategories(ctx context.Context) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockEventDBMockRecorder) GetCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockEventDB)(nil).GetCategories), ctx)
}

// GetEventByID mocks base method.
func (m *MockEventDB) GetEventByID(ctx context.Context, ID int) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventByID", ctx, ID)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventByID indicates an expected call of GetEventByID.
func (mr *MockEventDBMockRecorder) GetEventByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventByID", reflect.TypeOf((*MockEventDB)(nil).GetEventByID), ctx, ID)
}

// GetEventsByCategory mocks base method.
func (m *MockEventDB) GetEventsByCategory(ctx context.Context, categoryID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsByCategory", ctx, categoryID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsByCategory indicates an expected call of GetEventsByCategory.
func (mr *MockEventDBMockRecorder) GetEventsByCategory(ctx, categoryID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsByCategory", reflect.TypeOf((*MockEventDB)(nil).GetEventsByCategory), ctx, categoryID, paginationParams)
}

// GetEventsByUser mocks base method.
func (m *MockEventDB) GetEventsByUser(ctx context.Context, userID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsByUser", ctx, userID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsByUser indicates an expected call of GetEventsByUser.
func (mr *MockEventDBMockRecorder) GetEventsByUser(ctx, userID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsByUser", reflect.TypeOf((*MockEventDB)(nil).GetEventsByUser), ctx, userID, paginationParams)
}

// GetPastEvents mocks base method.
func (m *MockEventDB) GetPastEvents(ctx context.Context, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPastEvents", ctx, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPastEvents indicates an expected call of GetPastEvents.
func (mr *MockEventDBMockRecorder) GetPastEvents(ctx, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPastEvents", reflect.TypeOf((*MockEventDB)(nil).GetPastEvents), ctx, paginationParams)
}

// GetUpcomingEvents mocks base method.
func (m *MockEventDB) GetUpcomingEvents(ctx context.Context, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingEvents", ctx, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingEvents indicates an expected call of GetUpcomingEvents.
func (mr *MockEventDBMockRecorder) GetUpcomingEvents(ctx, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingEvents", reflect.TypeOf((*MockEventDB)(nil).GetUpcomingEvents), ctx, paginationParams)
}

// SearchEvents mocks base method.
func (m *MockEventDB) SearchEvents(ctx context.Context, params models.SearchParams, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchEvents", ctx, params, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEvents indicates an expected call of SearchEvents.
func (mr *MockEventDBMockRecorder) SearchEvents(ctx, params, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEvents", reflect.TypeOf((*MockEventDB)(nil).SearchEvents), ctx, params, paginationParams)
}

// UpdateEvent mocks base method.
func (m *MockEventDB) UpdateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", ctx, event)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockEventDBMockRecorder) UpdateEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockEventDB)(nil).UpdateEvent), ctx, event)
}
