# Lobby Service

The Lobby Service manages game lobbies where users gather before starting games.

## Features

- Create, join, and leave lobbies
- List available lobbies
- Lobby chat functionality
- Manage lobby settings
- Transition from lobby to game

## API

The service provides a gRPC API defined in `proto/lobby.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/lobby.proto

# Build the service
go build -o lobby-service ./cmd/server
```

### Running

```bash
# Run directly
./lobby-service

# Run with Docker
docker build -t storywarz/lobby-service .
docker run -p 50052:50052 storywarz/lobby-service
``` 