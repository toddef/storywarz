# Matchmaking Service

The Matchmaking Service pairs players for games based on skill level, preferences, and availability.

## Features

- Quick match functionality
- Skill-based matchmaking
- Match queuing system
- Custom match criteria
- Match confirmation and acceptance

## API

The service provides a gRPC API defined in `proto/matchmaking.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/matchmaking.proto

# Build the service
go build -o matchmaking-service ./cmd/server
```

### Running

```bash
# Run directly
./matchmaking-service

# Run with Docker
docker build -t storywarz/matchmaking-service .
docker run -p 50053:50053 storywarz/matchmaking-service
``` 