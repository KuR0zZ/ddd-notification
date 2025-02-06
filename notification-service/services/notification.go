package services

import (
	"context"
	"notification-service/domain/notification/entity"
	"notification-service/domain/notification/repository"
	"notification-service/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	pb.UnimplementedNotificationServiceServer
	Repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{Repository: repository}
}

func (ns *Service) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	notification := &entity.Notification{
		Email:   request.Email,
		Message: request.Message,
		Type:    request.Type,
		IsSent:  false,
	}

	if err := ns.Repository.Create(notification); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.CreateResponse{
		Id:   notification.ID,
		Type: notification.Type,
	}

	return response, nil
}

func (ns *Service) GetNotSent(ctx context.Context, request *emptypb.Empty) (*pb.Notifications, error) {
	return nil, nil
}

func (ns *Service) UpdateNotificationStatus(ctx context.Context, request *pb.Notifications) (*pb.UpdateStatusToSentResponse, error) {
	return nil, nil
}
