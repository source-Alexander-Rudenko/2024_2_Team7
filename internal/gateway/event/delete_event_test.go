package events

import (
	"net/http"
	"net/http/httptest"
	"testing"

	pb "kudago/internal/event/api"
	"kudago/internal/event/grpc"
	"kudago/internal/gateway/event/mocks"
	"kudago/internal/gateway/utils"
	"kudago/internal/logger"
	"kudago/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEventHandler_DeleteEvent(t *testing.T) {
	t.Parallel()

	deleteEventRequest := &pb.DeleteEventRequest{
		EventID:  1,
		AuthorID: 1,
	}

	logger, _ := logger.NewLogger()

	tests := []struct {
		name      string
		req       *http.Request
		setupFunc func(ctrl *gomock.Controller) *EventHandler
		wantCode  int
	}{
		{
			name: "Успешное ",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/events/1", nil)
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
				session := models.Session{UserID: 1, Token: "valid_token"}
				ctx := utils.SetSessionInContext(req.Context(), session)
				return req.WithContext(ctx)
			}(),
			setupFunc: func(ctrl *gomock.Controller) *EventHandler {
				serviceMock := mocks.NewMockEventServiceClient(ctrl)

				serviceMock.EXPECT().DeleteEvent(gomock.Any(), deleteEventRequest).Return(nil, nil)

				return &EventHandler{
					EventService: serviceMock,
					logger:       logger,
				}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "Not found",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/events/1", nil)
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
				session := models.Session{UserID: 1, Token: "valid_token"}
				ctx := utils.SetSessionInContext(req.Context(), session)
				return req.WithContext(ctx)
			}(),
			setupFunc: func(ctrl *gomock.Controller) *EventHandler {
				serviceMock := mocks.NewMockEventServiceClient(ctrl)
				serviceMock.EXPECT().DeleteEvent(gomock.Any(), deleteEventRequest).Return(nil, status.Error(codes.NotFound, grpc.ErrEventNotFound))

				return &EventHandler{
					EventService: serviceMock,
					logger:       logger,
				}
			},
			wantCode: http.StatusConflict,
		},
		{
			name: "Access denied",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/events/1", nil)
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
				session := models.Session{UserID: 1, Token: "valid_token"}
				ctx := utils.SetSessionInContext(req.Context(), session)
				return req.WithContext(ctx)
			}(),
			setupFunc: func(ctrl *gomock.Controller) *EventHandler {
				serviceMock := mocks.NewMockEventServiceClient(ctrl)
				serviceMock.EXPECT().DeleteEvent(gomock.Any(), deleteEventRequest).Return(nil, status.Error(codes.PermissionDenied, grpc.ErrPermissionDenied))

				return &EventHandler{
					EventService: serviceMock,
					logger:       logger,
				}
			},
			wantCode: http.StatusForbidden,
		},
		{
			name: "Internal error",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodDelete, "/events/1", nil)
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
				session := models.Session{UserID: 1, Token: "valid_token"}
				ctx := utils.SetSessionInContext(req.Context(), session)
				return req.WithContext(ctx)
			}(),
			setupFunc: func(ctrl *gomock.Controller) *EventHandler {
				serviceMock := mocks.NewMockEventServiceClient(ctrl)
				serviceMock.EXPECT().DeleteEvent(gomock.Any(), deleteEventRequest).Return(nil, status.Error(codes.Internal, grpc.ErrInternal))

				return &EventHandler{
					EventService: serviceMock,
					logger:       logger,
				}
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			recorder := httptest.NewRecorder()
			tt.setupFunc(ctrl).DeleteEvent(recorder, tt.req)

			assert.Equal(t, tt.wantCode, recorder.Code)
		})
	}
}
