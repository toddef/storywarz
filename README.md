# StoryWarz

StoryWarz is a real-time multiplayer game where players compete to write the most creative and entertaining story snippets based on prompts.

## Project Structure

The project uses a microservice architecture with the following components:

- **User Service**: Manages user accounts, authentication, and profiles
- **Lobby Service**: Handles game lobbies and chat functionality
- **Matchmaking Service**: Pairs players for games based on preferences
- **Game Service**: Handles core gameplay mechanics and state
- **Scoring Service**: Manages voting and score calculations
- **Prompt Service**: Provides story prompts and categories
- **Notification Service**: Delivers real-time updates and notifications
- **GraphQL Gateway**: Provides a unified API for front-end clients

## Prerequisites

- Docker and Docker Compose
- Go 1.22+ (for local development)
- Protocol Buffer Compiler (for generating gRPC code)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/storywarz.git
   cd storywarz
   ```

2. Run the setup test script to verify your environment:
   ```bash
   # On Unix/Linux/Mac
   ./test_setup.sh
   
   # On Windows
   test_setup.bat
   ```

3. Start all services:
   ```bash
   docker-compose up
   ```

   Or run in detached mode:
   ```bash
   docker-compose up -d
   ```

4. Access the GraphQL Playground at http://localhost:8080/playground

## Development

Each service has its own directory with:
- gRPC protocol definitions in the `proto` directory
- Dockerfile for containerization
- README with service-specific information

## Database

The project uses:
- PostgreSQL for persistent storage
- Redis for caching and pub/sub messaging

Database migrations will be automatically applied when containers start up.

## Testing

Run individual service tests:
```bash
# Navigate to a service directory
cd services/user-service

# Run tests
go test ./...
```

## Documentation

For detailed information about the project architecture and services, see the [spec.md](spec.md) file.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
