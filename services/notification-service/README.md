# Notification Service

The Notification Service manages in-app notifications, real-time updates, and external communications like email.

## Features

- In-app notification management
- Real-time updates using WebSockets
- Email notifications
- Notification preferences
- Notification history and read status

## API

The service provides a gRPC API defined in `proto/notification.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/notification.proto

# Build the service
go build -o notification-service ./cmd/server
```

### Running

```bash
# Run directly
./notification-service

# Run with Docker
docker build -t storywarz/notification-service .
docker run -p 50057:50057 storywarz/notification-service
``` 