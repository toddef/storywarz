# Game Service

The Game Service manages the core gameplay mechanics, including rounds, turns, submissions, and game state.

## Features

- Game creation and management
- Round and turn management
- Player submission handling
- Game state transitions
- Game results recording

## API

The service provides a gRPC API defined in `proto/game.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/game.proto

# Build the service
go build -o game-service ./cmd/server
```

### Running

```bash
# Run directly
./game-service

# Run with Docker
docker build -t storywarz/game-service .
docker run -p 50054:50054 storywarz/game-service
``` 