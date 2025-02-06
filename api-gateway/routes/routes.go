package routes

import (
	"api-gateway/controllers"

	"github.com/gofiber/fiber/v3"
)

func InitRoute(app *fiber.App, nc controllers.NotificationController) {
	app.Post("/notifications", nc.Create)
}
