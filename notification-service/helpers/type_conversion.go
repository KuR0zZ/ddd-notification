package helpers

import (
	"notification-service/domain/notification/entity"
	"notification-service/proto/pb"
)

func ConvertEntityToPb(notificationsTmp []entity.Notification) []*pb.Notification {
	var notifications []*pb.Notification

	for _, notificationTmp := range notificationsTmp {
		notification := &pb.Notification{
			Id:        notificationTmp.ID,
			Email:     notificationTmp.Email,
			Message:   notificationTmp.Message,
			Type:      notificationTmp.Type,
			IsSent:    notificationTmp.IsSent,
			CreatedAt: notificationTmp.CreatedAt.String(),
			UpdatedAt: notificationTmp.UpdatedAt.String(),
		}

		notifications = append(notifications, notification)
	}

	return notifications
}
