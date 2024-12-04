// Code generated by MockGen. DO NOT EDIT.
// Source: event.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "kudago/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// AddEvent mocks base method.
func (m *MockEventService) AddEvent(ctx context.Context, event models.Event) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEvent", ctx, event)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEvent indicates an expected call of AddEvent.
func (mr *MockEventServiceMockRecorder) AddEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEvent", reflect.TypeOf((*MockEventService)(nil).AddEvent), ctx, event)
}

// AddEventToFavorites mocks base method.
func (m *MockEventService) AddEventToFavorites(ctx context.Context, newFavorite models.FavoriteEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventToFavorites", ctx, newFavorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventToFavorites indicates an expected call of AddEventToFavorites.
func (mr *MockEventServiceMockRecorder) AddEventToFavorites(ctx, newFavorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventToFavorites", reflect.TypeOf((*MockEventService)(nil).AddEventToFavorites), ctx, newFavorite)
}

// DeleteEvent mocks base method.
func (m *MockEventService) DeleteEvent(ctx context.Context, ID, authorID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", ctx, ID, authorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockEventServiceMockRecorder) DeleteEvent(ctx, ID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockEventService)(nil).DeleteEvent), ctx, ID, authorID)
}

// DeleteEventFromFavorites mocks base method.
func (m *MockEventService) DeleteEventFromFavorites(ctx context.Context, newFavorite models.FavoriteEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventFromFavorites", ctx, newFavorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventFromFavorites indicates an expected call of DeleteEventFromFavorites.
func (mr *MockEventServiceMockRecorder) DeleteEventFromFavorites(ctx, newFavorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventFromFavorites", reflect.TypeOf((*MockEventService)(nil).DeleteEventFromFavorites), ctx, newFavorite)
}

// SearchEvents mocks base method.
func (m *MockEventService) SearchEvents(ctx context.Context, params models.SearchParams, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchEvents", ctx, params, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEvents indicates an expected call of SearchEvents.
func (mr *MockEventServiceMockRecorder) SearchEvents(ctx, params, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEvents", reflect.TypeOf((*MockEventService)(nil).SearchEvents), ctx, params, paginationParams)
}

// UpdateEvent mocks base method.
func (m *MockEventService) UpdateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", ctx, event)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockEventServiceMockRecorder) UpdateEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockEventService)(nil).UpdateEvent), ctx, event)
}

// MockEventsGetter is a mock of EventsGetter interface.
type MockEventsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEventsGetterMockRecorder
}

// MockEventsGetterMockRecorder is the mock recorder for MockEventsGetter.
type MockEventsGetterMockRecorder struct {
	mock *MockEventsGetter
}

// NewMockEventsGetter creates a new mock instance.
func NewMockEventsGetter(ctrl *gomock.Controller) *MockEventsGetter {
	mock := &MockEventsGetter{ctrl: ctrl}
	mock.recorder = &MockEventsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventsGetter) EXPECT() *MockEventsGetterMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockEventsGetter) GetCategories(ctx context.Context) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockEventsGetterMockRecorder) GetCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockEventsGetter)(nil).GetCategories), ctx)
}

// GetEventByID mocks base method.
func (m *MockEventsGetter) GetEventByID(ctx context.Context, ID int) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventByID", ctx, ID)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventByID indicates an expected call of GetEventByID.
func (mr *MockEventsGetterMockRecorder) GetEventByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventByID", reflect.TypeOf((*MockEventsGetter)(nil).GetEventByID), ctx, ID)
}

// GetEventsByCategory mocks base method.
func (m *MockEventsGetter) GetEventsByCategory(ctx context.Context, categoryID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsByCategory", ctx, categoryID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsByCategory indicates an expected call of GetEventsByCategory.
func (mr *MockEventsGetterMockRecorder) GetEventsByCategory(ctx, categoryID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsByCategory", reflect.TypeOf((*MockEventsGetter)(nil).GetEventsByCategory), ctx, categoryID, paginationParams)
}

// GetEventsByUser mocks base method.
func (m *MockEventsGetter) GetEventsByUser(ctx context.Context, userID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsByUser", ctx, userID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsByUser indicates an expected call of GetEventsByUser.
func (mr *MockEventsGetterMockRecorder) GetEventsByUser(ctx, userID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsByUser", reflect.TypeOf((*MockEventsGetter)(nil).GetEventsByUser), ctx, userID, paginationParams)
}

// GetFavorites mocks base method.
func (m *MockEventsGetter) GetFavorites(ctx context.Context, userID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", ctx, userID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockEventsGetterMockRecorder) GetFavorites(ctx, userID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockEventsGetter)(nil).GetFavorites), ctx, userID, paginationParams)
}

// GetPastEvents mocks base method.
func (m *MockEventsGetter) GetPastEvents(ctx context.Context, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPastEvents", ctx, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPastEvents indicates an expected call of GetPastEvents.
func (mr *MockEventsGetterMockRecorder) GetPastEvents(ctx, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPastEvents", reflect.TypeOf((*MockEventsGetter)(nil).GetPastEvents), ctx, paginationParams)
}

// GetSubscriptionEvents mocks base method.
func (m *MockEventsGetter) GetSubscriptionEvents(ctx context.Context, userID int, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionEvents", ctx, userID, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionEvents indicates an expected call of GetSubscriptionEvents.
func (mr *MockEventsGetterMockRecorder) GetSubscriptionEvents(ctx, userID, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionEvents", reflect.TypeOf((*MockEventsGetter)(nil).GetSubscriptionEvents), ctx, userID, paginationParams)
}

// GetUpcomingEvents mocks base method.
func (m *MockEventsGetter) GetUpcomingEvents(ctx context.Context, paginationParams models.PaginationParams) ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingEvents", ctx, paginationParams)
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingEvents indicates an expected call of GetUpcomingEvents.
func (mr *MockEventsGetterMockRecorder) GetUpcomingEvents(ctx, paginationParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingEvents", reflect.TypeOf((*MockEventsGetter)(nil).GetUpcomingEvents), ctx, paginationParams)
}