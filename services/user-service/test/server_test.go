package test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/storywarz/user-service/internal/server"
	pb "github.com/storywarz/user-service/pkg/proto"
)

const (
	testAddress = "127.0.0.1"
	testPort    = 8002 // Use a different port for testing
)

func TestServerStartsAndResponds(t *testing.T) {
	// Create a logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Create and start the server in a separate goroutine
	srv := server.NewServer(testAddress, testPort, logger)
	go func() {
		if err := srv.Start(); err != nil {
			// Only fail if it's not a "use of closed network connection" error,
			// which can happen during shutdown
			if err.Error() != "accept tcp 127.0.0.1:8002: use of closed network connection" {
				t.Errorf("Failed to start server: %v", err)
			}
		}
	}()

	// Give the server time to start
	time.Sleep(1 * time.Second)

	// Set up a connection to the server
	conn, err := grpc.Dial(
		net.JoinHostPort(testAddress, "8002"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewUserServiceClient(conn)

	// Call GetUser
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// Check the response
	if resp.User == nil {
		t.Fatal("Expected user in response, got nil")
	}
	if resp.User.Id != "123" {
		t.Errorf("Expected user ID 123, got %s", resp.User.Id)
	}

	// Stop the server
	srv.Stop()
}
