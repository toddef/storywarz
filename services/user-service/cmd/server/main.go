package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/storywarz/user-service/internal/server"
)

const (
	serverAddress = "0.0.0.0"
	serverPort    = 8001
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// Create and start the gRPC server
	grpcServer := server.NewServer(serverAddress, serverPort, logger)

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Start(); err != nil {
			logger.WithError(err).Fatal("Failed to start gRPC server")
		}
	}()

	logger.Info("User service started")

	// Wait for termination signal
	sig := <-sigCh
	logger.WithField("signal", sig.String()).Info("Received termination signal")

	// Stop the server
	grpcServer.Stop()
	logger.Info("User service stopped")
}
