package services

import (
	"ddd-notification/domain/notification"
	"ddd-notification/domain/notification/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotificationService struct {
	repo notification.Repository
}

func NewNotificationService(repo notification.Repository) *NotificationService {
	return &NotificationService{repo}
}

func (ns *NotificationService) AddNotification(c echo.Context) error {
	var req struct {
		Email   string `json:"email" validate:"required,email"`
		Message string `json:"message" validate:"required"`
		Type    string `json:"type" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	if req.Type != "SMS" && req.Type != "EMAIL" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid notification type")
	}

	data := &entity.Notification{
		Email:   req.Email,
		Message: req.Message,
		Type:    req.Type,
		Status:  "Not Sent",
	}

	if err := ns.repo.Add(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Notification added successfully",
		"data":    data,
	})
}

func (ns *NotificationService) GetNotSentNotification(c echo.Context) error {
	notifications, err := ns.repo.GetNotSent()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, notifications)
}
