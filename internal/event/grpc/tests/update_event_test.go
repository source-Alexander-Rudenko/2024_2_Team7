package grpc

import (
	"context"
	"testing"

	pb "kudago/internal/event/api"
	"kudago/internal/event/grpc/tests/mocks"
	"kudago/internal/logger"
	"kudago/internal/models"

	event "kudago/internal/event/grpc"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEventGRPC_UpdateEvent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventData := models.Event{
		ID:    1,
		Title: "test",
	}

	tests := []struct {
		name         string
		req          *pb.Event
		setupFunc    func(ctrl *gomock.Controller) *event.ServerAPI
		expectedResp *pb.Event
		expectedErr  error
	}{
		{
			name: "success update event",
			req: &pb.Event{
				ID:    1,
				Title: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventService.EXPECT().
					UpdateEvent(context.Background(), eventData).
					Return(eventData, nil)
				return event.NewServerAPI(mockEventService, mockEventGetter, logger)
			},
			expectedResp: &pb.Event{
				ID:    1,
				Title: "test",
			},
			expectedErr: nil,
		},
		{
			name: "event not found",
			req: &pb.Event{
				ID:    1,
				Title: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventService.EXPECT().
					UpdateEvent(context.Background(), eventData).
					Return(models.Event{}, models.ErrEventNotFound)
				return event.NewServerAPI(mockEventService, mockEventGetter, logger)
			},
			expectedErr: status.Error(codes.NotFound, event.ErrEventNotFound),
		},
		{
			name: "access denied",
			req: &pb.Event{
				ID:    1,
				Title: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventService.EXPECT().
					UpdateEvent(context.Background(), eventData).
					Return(models.Event{}, models.ErrAccessDenied)
				return event.NewServerAPI(mockEventService, mockEventGetter, logger)
			},
			expectedErr: status.Error(codes.PermissionDenied, event.ErrPermissionDenied),
		},
		{
			name: "internal error",
			req: &pb.Event{
				ID:    1,
				Title: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventService.EXPECT().
					UpdateEvent(context.Background(), eventData).
					Return(models.Event{}, models.ErrInternal)
				return event.NewServerAPI(mockEventService, mockEventGetter, logger)
			},
			expectedErr: status.Error(codes.Internal, event.ErrInternal),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			_, err := tt.setupFunc(ctrl).UpdateEvent(context.Background(), tt.req)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
