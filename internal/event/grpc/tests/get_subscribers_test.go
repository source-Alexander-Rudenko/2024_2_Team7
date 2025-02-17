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

func TestEventGRPC_GetSubscribers(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	subscribers := []int{1, 2}

	tests := []struct {
		name         string
		req          *pb.GetSubscribersIDsRequest
		setupFunc    func(ctrl *gomock.Controller) *event.ServerAPI
		expectedResp *pb.GetUserIDsResponse
		expectedErr  error
	}{
		{
			name: "success get subscribers",
			req: &pb.GetSubscribersIDsRequest{
				UserID: 1,
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventGetter.EXPECT().
					GetSubscribersIDs(context.Background(), 1).
					Return(subscribers, nil)
				return event.NewServerAPI(mockEventService, mockEventGetter, logger)
			},
			expectedResp: &pb.GetUserIDsResponse{
				IDs: []int32{1, 2},
			},
			expectedErr: nil,
		},
		{
			name: "internal error",
			req: &pb.GetSubscribersIDsRequest{
				UserID: 1,
			},
			setupFunc: func(ctrl *gomock.Controller) *event.ServerAPI {
				mockEventService := mocks.NewMockEventService(ctrl)
				mockEventGetter := mocks.NewMockEventsGetter(ctrl)

				logger, _ := logger.NewLogger()

				mockEventGetter.EXPECT().
					GetSubscribersIDs(context.Background(), 1).
					Return(nil, models.ErrInternal)
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

			_, err := tt.setupFunc(ctrl).GetSubscribersIDs(context.Background(), tt.req)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
