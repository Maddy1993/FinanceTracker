package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"gorm.io/gorm"
	pb "myapp/financetracker/internal/models"
)

// userServiceServer implements pb.UserServiceServer
type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

// NewUserServiceServer constructs a new service with a given GORM DB connection
func NewUserServiceServer(db *gorm.DB) pb.UserServiceServer {
	return &userServiceServer{
		db: db,
	}
}

// CreateUser inserts a new user into the Postgres DB
func (s *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Validate request
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetFullName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	// Build user record
	newUser := UserModel{
		Email: req.GetEmail(),
		Name:  req.GetFullName(),
	}

	// Insert into DB
	if err := s.db.Create(&newUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	log.Printf("Created user: %v", newUser)

	// Convert to protobuf User
	pbUser := &pb.User{
		UserId:   newUser.ID,
		Email:    newUser.Email,
		FullName: newUser.Name,
	}

	return &pb.CreateUserResponse{
		User: pbUser,
	}, nil
}
