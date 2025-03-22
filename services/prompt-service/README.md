# Prompt Service

The Prompt Service manages story prompts used in games, including prompt creation, categorization, and selection.

## Features

- Prompt management and storage
- Prompt categorization
- Random prompt selection
- User-submitted prompts
- Prompt review and approval

## API

The service provides a gRPC API defined in `proto/prompt.proto`.

## Development

### Requirements

- Go 1.22+
- Protocol Buffer Compiler
- Docker (for containerization)

### Building

```bash
# Generate Go code from protocol buffers
protoc --go_out=. --go-grpc_out=. proto/prompt.proto

# Build the service
go build -o prompt-service ./cmd/server
```

### Running

```bash
# Run directly
./prompt-service

# Run with Docker
docker build -t storywarz/prompt-service .
docker run -p 50056:50056 storywarz/prompt-service
``` 