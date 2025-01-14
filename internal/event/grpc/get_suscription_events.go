package grpc

import (
	"context"

	pb "kudago/internal/event/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) GetSubscriptionsEvents(ctx context.Context, req *pb.GetSubscriptionsRequest) (*pb.Events, error) {
	params := getPaginationParams(req.Params)
	eventsData, err := s.getter.GetSubscriptionEvents(ctx, int(req.ID), params)
	if err != nil {
		s.logger.Error(ctx, "add event", err)
		return nil, status.Error(codes.Internal, errInternal)
	}

	event := writeEventsResponse(eventsData, params.Limit)

	return event, nil
}
