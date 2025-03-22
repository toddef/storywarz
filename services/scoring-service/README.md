# Scoring Service

The Scoring Service manages the voting and scoring mechanics for game entries, including player votes, round scores, and overall game results.

## Features

- Vote submission and tallying
- Score calculation for rounds and games
- Leaderboard management
- Score history tracking
- Rating calculation

## API

The service provides a gRPC API defined in `proto/scoring.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/scoring.proto

# Build the service
go build -o scoring-service ./cmd/server
```

### Running

```bash
# Run directly
./scoring-service

# Run with Docker
docker build -t storywarz/scoring-service .
docker run -p 50055:50055 storywarz/scoring-service
``` 