# GraphQL Gateway

The GraphQL Gateway provides a unified API layer over all microservices, exposing a cohesive GraphQL API to clients.

## Features

- Single GraphQL endpoint for all services
- Authentication and authorization
- Request validation
- Rate limiting
- Error handling and formatting
- Comprehensive schema with documentation

## Development

### Requirements

- Go 1.22+
- gqlgen (GraphQL code generator)
- Docker (for containerization)

### Building

```bash
# Generate Go code from GraphQL schema
go run github.com/99designs/gqlgen generate

# Build the service
go build -o graphql-gateway ./cmd/server
```

### Running

```bash
# Run directly
./graphql-gateway

# Run with Docker
docker build -t storywarz/graphql-gateway .
docker run -p 8080:8080 storywarz/graphql-gateway
```

### GraphQL Playground

When running in development mode, a GraphQL Playground is available at http://localhost:8080/playground for interactive exploration of the API. 