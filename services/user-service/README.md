# User Service

The User Service manages user-related operations like registration, authentication, profile management, and user preferences.

## Features

- User registration and authentication
- User profile management
- Password reset functionality
- Session management
- User preferences

## API

The service provides a gRPC API defined in `proto/user.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/user.proto

# Build the service
go build -o user-service ./cmd/server
```

### Running

```bash
# Run directly
./user-service

# Run with Docker
docker build -t storywarz/user-service .
docker run -p 50051:50051 storywarz/user-service
``` 