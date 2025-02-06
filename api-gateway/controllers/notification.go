package controllers

import (
	"api-gateway/dtos"
	"api-gateway/helpers"
	"api-gateway/proto/pb"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc/status"
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
