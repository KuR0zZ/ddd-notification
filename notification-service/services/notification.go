package services

import (
	"context"
	"errors"
	"notification-service/domain/notification/entity"
	"notification-service/domain/notification/repository"
	"notification-service/helpers"
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

func (s *Service) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	notification := &entity.Notification{
		Email:   request.Email,
		Message: request.Message,
		Type:    request.Type,
		IsSent:  false,
	}

	if err := s.Repository.Create(notification); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.CreateResponse{
		Id:   notification.ID,
		Type: notification.Type,
	}

	return response, nil
}

func (s *Service) GetNotSent(ctx context.Context, request *emptypb.Empty) (*pb.Notifications, error) {
	notificationsTmp, err := s.Repository.GetNotSent()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	notifications := helpers.ConvertEntityToPb(notificationsTmp)

	return &pb.Notifications{Notification: notifications}, nil
}

func (s *Service) UpdateStatusToSent(ctx context.Context, request *pb.Notifications) (*pb.UpdateStatusToSentResponse, error) {
	var email_counter uint32
	var sms_counter uint32

	for _, notification := range request.Notification {
		if err := s.Repository.UpdateStatus(notification.Id); err != nil {
			if errors.Is(err, repository.ErrNotificationNotFound) {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, status.Error(codes.Internal, err.Error())
		}

		if notification.Type == "EMAIL" {
			email_counter++
		} else {
			sms_counter++
		}
	}

	return &pb.UpdateStatusToSentResponse{
		EmailNotification: email_counter,
		SmsNotification:   sms_counter,
	}, nil
}
