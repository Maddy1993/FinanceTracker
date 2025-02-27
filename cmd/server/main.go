package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"myapp/financetracker/internal/db"
	pb "myapp/financetracker/internal/models"
	"myapp/financetracker/internal/service"
)

func main() {
	// 1. Connect to Postgres
	database, err := db.ConnectPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	// 2. Auto-migrate the UserModel (creates "users" table if not exists)
	if err := database.Conn.AutoMigrate(&service.UserModel{}); err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	// 3. Setup gRPC server
	grpcServer := grpc.NewServer()

	// 4. Register our service
	userService := service.NewUserServiceServer(database.Conn)
	pb.RegisterUserServiceServer(grpcServer, userService)

	// 5. Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on :50051: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
