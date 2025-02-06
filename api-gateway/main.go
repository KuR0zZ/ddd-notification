package main

import (
	"api-gateway/controllers"
	custom_middleware "api-gateway/middleware"
	"api-gateway/proto/pb"
	"api-gateway/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: &custom_middleware.StructValidator{Validator: validator.New()},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	conn, err := grpc.NewClient(os.Getenv("NOTIFICATION_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	client := pb.NewNotificationServiceClient(conn)

	notificationController := controllers.NewNotificationController(client)

	routes.InitRoute(app, *notificationController)

	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	port := os.Getenv("PORT")
	go func() {
		if err := app.Listen(":" + port); err != nil {
			errChan <- err
		}
	}()

	defer func() {
		log.Println("Closing gRPC client connection...")
		conn.Close()

		log.Println("Shutting down Fiber server...")
		app.Shutdown()
	}()

	select {
	case err := <-errChan:
		log.Printf("Fail to serve: %v\n", err)
	case <-stopChan:
		log.Println("Received shutdown signal...")
	}
}
