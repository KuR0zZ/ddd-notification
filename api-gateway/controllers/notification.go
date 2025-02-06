package controllers

import (
	"api-gateway/dtos"
	"api-gateway/helpers"
	"api-gateway/proto/pb"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationController struct {
	Client pb.NotificationServiceClient
}

func NewNotificationController(client pb.NotificationServiceClient) *NotificationController {
	return &NotificationController{Client: client}
}

func (nc *NotificationController) Create(c fiber.Ctx) error {
	request := new(dtos.CreateRequest)

	if err := c.Bind().JSON(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if request.Type != "EMAIL" && request.Type != "SMS" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid notification type",
		})
	}

	data := &pb.CreateRequest{
		Email:   request.Email,
		Message: request.Message,
		Type:    request.Type,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer cancel()

	resGrpc, err := nc.Client.Create(ctx, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": status.Convert(err).Message(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully created notification with Id: %s and Type: %s", resGrpc.Id, resGrpc.Type),
	})
}

func (nc *NotificationController) GetNotSent(c fiber.Ctx) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer cancel()

	resGrpc, err := nc.Client.GetNotSent(ctx, &emptypb.Empty{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": status.Convert(err).Message(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resGrpc.Notification)
}

func (nc *NotificationController) UpdateStatusToSent(c fiber.Ctx) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer cancel()

	resGrpc1, err := nc.Client.GetNotSent(ctx, &emptypb.Empty{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": status.Convert(err).Message(),
		})
	}

	ctx, cancel, err = helpers.NewServiceContext()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer cancel()

	resGrpc2, err := nc.Client.UpdateStatusToSent(ctx, resGrpc1)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": err.Error(),
				})
			case codes.Internal:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully updated %d email notification and %d sms notification", resGrpc2.EmailNotification, resGrpc2.SmsNotification),
	})
}
