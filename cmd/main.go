package main

import (
	"ddd-notification/config"
	"ddd-notification/domain/notification"
	custom_middleware "ddd-notification/middleware"
	"ddd-notification/routes"
	"ddd-notification/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.InitDB()

	repository := notification.NewPostgresRepository(db)
	service := services.NewNotificationService(repository)

	e := echo.New()
	e.Validator = custom_middleware.NewValidate(validator.New())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoute(e, *service)

	e.Logger.Fatal(e.Start(":8080"))
}
