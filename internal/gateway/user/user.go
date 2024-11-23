package handlers

import (
	"kudago/internal/logger"
	pb "kudago/internal/user/api"
	user "kudago/internal/user/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserHandlers struct {
	UserService pb.UserServiceClient
	logger      *logger.Logger
}

func NewUserHandlers(userServiceAddr string, logger *logger.Logger) (*UserHandlers, error) {
	authConn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &UserHandlers{
		UserService: user.NewUserServiceClient(authConn),
		logger:      logger,
	}, nil
}

type AuthResponse struct {
	User UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ImageURL string `json:"image"`
}

type GetUsersResponse struct {
	Users []UserResponse `json:"users"`
}

func userToUserResponse(user *pb.User) UserResponse {
	return UserResponse{
		ID:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		// ImageURL: user.ImageURL,
	}
}

func subscriptionToSubscriptionPB(subscription *pb.Subscription) pb.Subscription {
	return pb.Subscription{
		SubscriberID: int32(subscription.SubscriberID),
		FollowsID:    int32(subscription.SubscriberID),
	}
}

func writeUsersResponse(users []*pb.User, limit int) GetUsersResponse {
	resp := GetUsersResponse{Users: make([]UserResponse, 0, limit)}

	for _, user := range users {
		userResp := userToUserResponse(user)
		resp.Users = append(resp.Users, userResp)
	}
	return resp
}