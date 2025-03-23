# User Service

A gRPC service for managing users in the StoryWarz platform.

## Features

- User management (create, read, update, delete)
- Authentication and authorization (login, registration, password reset)
- Profile management

## Technology Stack

- Go 1.22+
- gRPC
- Protocol Buffers
- PostgreSQL (for data persistence)
- Redis (for caching and sessions)

## API

The service exposes the following gRPC endpoints:

- `GetUser`: Retrieve a user by ID
- `CreateUser`: Register a new user
- `UpdateUser`: Update user profile information
- `DeleteUser`: Remove a user account

Check the `proto/user.proto` file for the complete API definition.

## Development

### Prerequisites

- Go 1.22+
- Protocol Buffer Compiler (`protoc`)
  - [Installation instructions](https://grpc.io/docs/protoc-installation/)
- Go plugins for Protocol Buffers:
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
- PostgreSQL (for production use)
- Redis (for production use)

### Setup

1. Clone the repository
2. Generate Protocol Buffer code:
   - Using make: `make proto`
   - Using scripts:
     - Windows: `scripts\generate_proto.bat`
     - Unix/Linux: `./scripts/generate_proto.sh`
3. Build the service:
   - Using make: `make build`
   - Manually: `go build -o bin/user-service ./cmd/server`
4. Run the service:
   - Using make: `make run`
   - Manually: `go run ./cmd/server/main.go`

### Testing

Run the tests:

```
make test
```

Or manually:

```
go test -v ./test/...
```

## Configuration

The service can be configured using environment variables:

- `PORT`: The port to listen on (default: 8001)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `REDIS_HOST`: Redis host
- `REDIS_PORT`: Redis port
- `REDIS_PASSWORD`: Redis password

## Docker

Build the Docker image:

```
docker build -t storywarz-user-service .
```

Run the container:

```
docker run -p 8001:8001 storywarz-user-service
``` 