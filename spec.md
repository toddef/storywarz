# StoryWarz Project Specification

## Overview

StoryWarz is a real-time multiplayer game where players compete to write the most creative and entertaining story snippets based on prompts. Players join lobbies, receive prompts, submit their entries, and vote on each other's submissions. Points are awarded based on votes, and winners are determined by their accumulated scores.

## Architecture

StoryWarz uses a microservice architecture with the following components:

1. **User Service**: Manages user accounts, authentication, and profiles
2. **Lobby Service**: Handles game lobbies and chat functionality
3. **Matchmaking Service**: Pairs players for games based on preferences
4. **Game Service**: Handles core gameplay mechanics and state
5. **Scoring Service**: Manages voting and score calculations
6. **Prompt Service**: Provides story prompts and categories
7. **Notification Service**: Delivers real-time updates and notifications
8. **GraphQL Gateway**: Provides a unified API for front-end clients

Services communicate via gRPC, with the GraphQL gateway serving as the primary entry point for clients.

## Data Store

- **PostgreSQL**: Primary relational database for structured data
- **Redis**: Caching, session management, and pub/sub for real-time features

## Service Details

### User Service (port 50051)

Responsible for:
- User registration and authentication
- JWT token generation and validation
- Profile management
- Password reset functionality
- Session management

### Lobby Service (port 50052)

Responsible for:
- Creating and managing game lobbies
- Player joining/leaving mechanics
- Lobby chat functionality
- Lobby settings configuration
- Transitioning from lobby to game

### Matchmaking Service (port 50053)

Responsible for:
- Quick match functionality
- Skill-based matchmaking
- Match queuing and pairing
- Custom match creation
- Match acceptance confirmation

### Game Service (port 50054)

Responsible for:
- Game state management
- Round and turn coordination
- Player submission handling
- Timer management
- Game results calculation

### Scoring Service (port 50055)

Responsible for:
- Vote collection and validation
- Score calculation
- Leaderboard management
- Player stats tracking
- Results aggregation

### Prompt Service (port 50056)

Responsible for:
- Prompt management and storage
- Prompt categorization
- Random prompt selection
- User-submitted prompts
- Prompt moderation

### Notification Service (port 50057)

Responsible for:
- In-app notifications
- Real-time updates via WebSockets
- Email notifications
- Push notifications
- Notification preferences

### GraphQL Gateway (port 8080)

Responsible for:
- Unified API layer for all services
- Authentication middleware
- Request validation
- Error handling
- WebSocket subscriptions for real-time updates

## Client Applications

The following client applications will interact with the GraphQL API:

1. **Web Application**: React-based web client
2. **Mobile Application**: React Native mobile app
3. **Admin Dashboard**: Management interface for administration

## Deployment

The application will be containerized using Docker and deployed with Docker Compose for development and Kubernetes for production.

## Development Workflow

1. Define service contracts using Protocol Buffers
2. Implement services in Go
3. Set up unit and integration tests
4. Build Docker images for each service
5. Deploy with Docker Compose for local development
6. Use CI/CD for automated testing and deployment

## Security Considerations

- All API endpoints must be authenticated except for registration and login
- Sensitive data must be encrypted at rest and in transit
- Rate limiting should be implemented to prevent abuse
- Input validation must be thorough
- Audit logging for security-sensitive operations

## Performance Goals

- API response times under 100ms for 95% of requests
- Support for 1000+ concurrent users
- Game sessions with up to 10 players
- Scalable architecture to handle traffic spikes

## Monitoring and Logging

- Distributed tracing across services
- Centralized logging
- Performance metrics collection
- Alerting for critical issues
- Health check endpoints for each service
