package routes

import (
	"ddd-notification/services"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, ns services.NotificationService) {
	e.POST("/notifications", ns.AddNotification)
}
