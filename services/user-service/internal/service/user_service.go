package service

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/storywarz/user-service/pkg/proto"
)

// UserService implements the gRPC UserService
type UserService struct {
	pb.UnimplementedUserServiceServer
	logger *logrus.Logger
}

// NewUserService creates a new UserService
func NewUserService(logger *logrus.Logger) *UserService {
	return &UserService{
		logger: logger,
	}
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.logger.WithField("user_id", req.Id).Info("GetUser called")

	// TODO: Implement database lookup

	// For now, just return a placeholder user
	user := &pb.User{
		Id:          req.Id,
		Username:    "user" + req.Id,
		Email:       "user" + req.Id + "@example.com",
		DisplayName: "User " + req.Id,
		CreatedAt:   1647209887,
		UpdatedAt:   1647209887,
	}

	return &pb.GetUserResponse{
		User: user,
	}, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.logger.WithFields(logrus.Fields{
		"username": req.Username,
		"email":    req.Email,
	}).Info("CreateUser called")

	// TODO: Implement user creation in database

	// For now, just return a placeholder user with a generated ID
	user := &pb.User{
		Id:          "user123",
		Username:    req.Username,
		Email:       req.Email,
		DisplayName: req.DisplayName,
		AvatarUrl:   req.AvatarUrl,
		CreatedAt:   1647209887,
		UpdatedAt:   1647209887,
	}

	return &pb.CreateUserResponse{
		User: user,
	}, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.logger.WithField("user_id", req.Id).Info("UpdateUser called")

	// TODO: Implement user update in database

	// For now, just return a placeholder updated user
	user := &pb.User{
		Id:          req.Id,
		Username:    "updated_user",
		Email:       "updated_user@example.com",
		DisplayName: "Updated User",
		CreatedAt:   1647209887,
		UpdatedAt:   1647210000,
	}

	return &pb.UpdateUserResponse{
		User: user,
	}, nil
}

// DeleteUser removes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.logger.WithField("user_id", req.Id).Info("DeleteUser called")

	// TODO: Implement user deletion in database

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
