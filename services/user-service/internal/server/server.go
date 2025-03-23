package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/storywarz/user-service/internal/service"
	pb "github.com/storywarz/user-service/pkg/proto"
)

// Server represents the gRPC server
type Server struct {
	address     string
	port        int
	logger      *logrus.Logger
	grpcServer  *grpc.Server
	userService *service.UserService
}

// NewServer creates a new gRPC server
func NewServer(address string, port int, logger *logrus.Logger) *Server {
	return &Server{
		address: address,
		port:    port,
		logger:  logger,
	}
}

// Start starts the gRPC server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.address, s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.grpcServer = grpc.NewServer()

	// Initialize and register the user service
	s.userService = service.NewUserService(s.logger)
	pb.RegisterUserServiceServer(s.grpcServer, s.userService)

	// Register the health service
	// healthService := health.NewServer()
	// grpc_health_v1.RegisterHealthServer(s.grpcServer, healthService)
	// healthService.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	// healthService.SetServingStatus("user.UserService", grpc_health_v1.HealthCheckResponse_SERVING)

	s.logger.WithField("address", fmt.Sprintf("%s:%d", s.address, s.port)).Info("Starting gRPC server")

	// Start the gRPC server
	return s.grpcServer.Serve(lis)
}

// Stop stops the gRPC server
func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.logger.Info("Stopping gRPC server")
		s.grpcServer.GracefulStop()
	}
}
