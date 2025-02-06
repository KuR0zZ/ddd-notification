package main

import (
	"log"
	"net"
	"notification-service/config"
	"notification-service/domain/notification/repository"
	"notification-service/middleware"
	"notification-service/proto/pb"
	"notification-service/services"
	"os"
	"os/signal"
	"syscall"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := config.InitDB()
	repository := repository.NewPostgresRepository(db)
	service := services.NewService(repository)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(middleware.JWTAuth),
		),
	)

	pb.RegisterNotificationServiceServer(grpcServer, service)

	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	log.Println("gRPC server is running on port:", port)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	defer func() {
		log.Println("Shutting down gRPC server...")
		grpcServer.GracefulStop()

		log.Println("Closing database connection...")
		db.Close()
	}()

	select {
	case err := <-errChan:
		log.Printf("Fail to serve: %v\n", err)
	case <-stopChan:
		log.Println("Received shutdown signal...")
	}
}
