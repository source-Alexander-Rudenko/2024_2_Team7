//go:generate mockgen -source=image.go -destination=tests/mocks/image.go -package=mocks

package grpc

import (
	"context"

	pb "kudago/internal/image/api"
	"kudago/internal/logger"
	"kudago/internal/models"
)

const (
	ErrInternal = "internal error"
)

type ServerAPI struct {
	pb.UnimplementedImageServiceServer
	service ImageService
	logger  *logger.Logger
}

type ImageService interface {
	UploadImage(ctx context.Context, media models.MediaFile) (string, error)
	DeleteImage(ctx context.Context, imagePath string) error
}

func NewServerAPI(service ImageService, logger *logger.Logger) *ServerAPI {
	return &ServerAPI{
		service: service,
		logger:  logger,
	}
}
