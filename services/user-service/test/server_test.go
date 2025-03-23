package test

import (
	"net"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/storywarz/user-service/internal/server"
)

const (
	testAddress = "127.0.0.1"
	testPort    = 8002 // Use a different port for testing
)

func TestServerStarts(t *testing.T) {
	// Create a logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Create the server
	srv := server.NewServer(testAddress, testPort, logger)

	// Start the server in a goroutine
	go func() {
		err := srv.Start()
		if err != nil {
			// We expect an error when the server is stopped
			t.Logf("Server stopped with error: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(500 * time.Millisecond)

	// Try to connect to the server to verify it's listening
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(testAddress, "8002"), 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}

	// If we get here, the connection was successful, so the server is listening
	t.Log("Successfully connected to server")

	// Close the connection
	conn.Close()

	// Stop the server
	srv.Stop()

	// Wait a moment for the server to fully stop
	time.Sleep(500 * time.Millisecond)
}
