===========================================================================
                           StoryWarz Specification
===========================================================================

1. OVERVIEW

StoryWarz is a multiplayer game where players join game lobbies using a unique join code.
In each game, players receive a prompt, submit several stories, and then guess which
story belongs to which player. The game is built using Golang 1.24 and follows a
microservices architecture organized as a monorepo.

===========================================================================
2. HIGH-LEVEL REQUIREMENTS

- Programming Language: Golang 1.24
- Architecture: Microservices in a monorepo with individual services for distinct responsibilities.
- Frontend & Backend:
  - Frontend: A simple, extensible web application.
  - Backend: Multiple services communicating via gRPC.
- Inter-Service Communication: gRPC with unary RPC calls (initially).
- API Gateway: GraphQL Gateway (using gqlgen) to aggregate data from backend services.
- Data Persistence: PostgreSQL for persistent data.
- Caching: Redis for caching active states and quick lookups.
- Authentication: Google OAuth for user authentication.
- Containerization & Deployment: Each service has its own Dockerfile.
  A top-level Docker Compose file orchestrates PostgreSQL, Redis, and all services.

===========================================================================
3. ARCHITECTURAL OVERVIEW

3.1 Monorepo Structure

   /storywarz
     ├── docker-compose.yaml
     ├── spec.md            (This file)
     ├── services/
     │     ├── user-service/
     │     │      ├── proto/user.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     ├── lobby-service/
     │     │      ├── proto/lobby.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     ├── matchmaking-service/
     │     │      ├── proto/matchmaking.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     ├── game-service/
     │     │      ├── proto/game.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     ├── scoring-service/
     │     │      ├── proto/scoring.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     ├── prompt-service/
     │     │      ├── proto/prompt.proto
     │     │      ├── Dockerfile
     │     │      └── README.md
     │     └── notification-service/
     │            ├── proto/notification.proto
     │            ├── Dockerfile
     │            └── README.md
     └── graphql-gateway/
           ├── schema.graphql
           ├── resolvers.go
           ├── gqlgen.yml
           ├── Dockerfile
           └── README.md

3.2 Backend Services and Responsibilities

  1. User Service (user-service)
     - Responsibilities: User account management, authentication (OAuth callback), JWT token generation.
     - Protobuf: proto/user.proto (methods: GetUser, CreateOrUpdateUser)

  2. Lobby Service (lobby-service)
     - Responsibilities: Creation and management of game lobbies, join-by-code functionality.
     - Protobuf: proto/lobby.proto (methods: CreateLobby, JoinLobby, GetLobby)

  3. Matchmaking Service (matchmaking-service)
     - Responsibilities: Listing public lobbies and auto-joining players.
     - Protobuf: proto/matchmaking.proto

  4. Game Management Service (game-service)
     - Responsibilities: Handling game state transitions, rounds, story submissions, and guesses.
     - Protobuf: proto/game.proto (methods: StartGame, SubmitStory, SubmitGuess, GetGame)

  5. Scoring Service (scoring-service)
     - Responsibilities: Calculating and tracking player scores.
     - Protobuf: proto/scoring.proto

  6. Prompt Service (prompt-service)
     - Responsibilities: Managing and delivering theme prompts.
     - Protobuf: proto/prompt.proto

  7. Notification Service (notification-service)
     - Responsibilities: Real-time notifications and event dispatch (e.g., game state changes).
     - Protobuf: proto/notification.proto

  8. API Gateway (GraphQL Gateway)
     - Responsibilities: Aggregates data from backend services and exposes a unified GraphQL API.
     - Technology: Built with gqlgen.
     - Communication: Uses gRPC clients to interact with backend services.

===========================================================================
4. DATA HANDLING

4.1 Database

  - Persistent Storage: PostgreSQL is used for persistent data (users, lobbies, games, scores, prompts, etc.).
  - Example User Schema:
    -----------------------------------------------------------
    CREATE TABLE users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        email VARCHAR(255) UNIQUE NOT NULL,
        name VARCHAR(255),
        profile_picture TEXT,
        oauth_provider VARCHAR(50) NOT NULL,
        oauth_provider_id VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );
    -----------------------------------------------------------
  - Caching: Redis is used to cache active lobby codes, temporary game states, session info, and other frequently accessed data.

4.2 gRPC & Protobufs

  - Communication: Each service exposes a gRPC interface using unary RPC calls.
  - Protobuf Generation: Use the protoc compiler with Go plugins (protoc-gen-go, protoc-gen-go-grpc).
    Example command:
      protoc --go_out=. --go-grpc_out=. services/user-service/proto/user.proto
  - Inter-Service Calls: The GraphQL Gateway initializes gRPC clients for each service and uses them in resolvers.

===========================================================================
5. GRAPHQL API DESIGN

5.1 Schema Considerations

  - Modular Schema: Organize the schema by domain (User, Lobby, Game, etc.).
  - Custom Scalars: Define a DateTime scalar for timestamps.
  - Input Types: Use input objects for mutations (e.g., CreateLobbyInput).
  - Sample Schema Snippet:
    ---------------------------------------------------------------------
    """User details from OAuth login."""
    type User {
      id: ID!
      email: String!
      name: String!
      profilePicture: String
    }

    """Represents a game lobby."""
    type Lobby {
      id: ID!
      joinCode: String!
      status: String!
      players: [User!]!
      isPublic: Boolean!
      createdAt: DateTime!
    }

    scalar DateTime

    type Query {
      """Fetches the authenticated user's profile."""
      me: User!
      """Fetch lobby details by lobby ID."""
      lobby(lobbyId: ID!): Lobby
    }

    input CreateLobbyInput {
      isPublic: Boolean!
    }

    type Mutation {
      """Creates a new game lobby."""
      createLobby(input: CreateLobbyInput!): Lobby!
      """Join an existing lobby by code."""
      joinLobby(joinCode: String!): Lobby!
    }
    ---------------------------------------------------------------------

5.2 Resolver Strategy

  - gRPC Integration: Resolvers call backend services via gRPC. Example: the 'me' query resolver calls UserService.GetUser.
  - Error Handling: Errors from gRPC calls are mapped to GraphQL error responses.
  - N+1 Problem Prevention: Use DataLoader patterns to batch nested queries (e.g., fetching players for a lobby).

===========================================================================
6. ERROR HANDLING STRATEGIES

- gRPC Timeout: Use context timeouts (e.g., 5 seconds) for all gRPC calls to avoid hanging requests.
- Error Mapping: Translate gRPC errors into meaningful GraphQL error messages.
- Logging & Monitoring: Use gRPC interceptors for logging and error management (future enhancement: add security).
- Standardized Responses: Maintain consistent error messages and codes across services.

===========================================================================
7. TESTING PLAN

7.1 Unit Testing

  - Services: Write unit tests for each service's business logic and gRPC method implementations.
  - GraphQL Resolvers: Test resolver logic, including mapping of gRPC responses to GraphQL types.
  - Tools: Utilize Go’s testing package and mocking frameworks to simulate gRPC calls.

7.2 Integration Testing

  - Inter-Service Communication: Validate that gRPC calls between services function as expected.
  - GraphQL Gateway: Test the full request flow from GraphQL query/mutation to backend gRPC response.
  - Environment: Use Docker Compose to spin up PostgreSQL, Redis, and all services for integration tests.

7.3 End-to-End Testing

  - User Journeys: Automate tests for key user workflows (e.g., login via Google OAuth, lobby creation/joining, game progression).
  - Tools: Consider using Postman, Cypress, or similar frameworks for frontend e2e tests.

7.4 Continuous Integration

  - CI Pipeline: Set up a CI pipeline to run all tests (unit, integration, e2e) on every commit.
  - Code Coverage: Monitor test coverage and ensure that critical paths are thoroughly tested.

===========================================================================
8. DEPLOYMENT

8.1 Docker Compose Setup

A top-level docker-compose.yaml orchestrates local development:

-----------------------------------------------------------
version: "3.9"
services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: storywarz_user
      POSTGRES_PASSWORD: secure_password
      POSTGRES_DB: storywarz_db
    ports:
      - "5432:5432"
    volumes:
      - storywarz_postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  user-service:
    build:
      context: ./services/user-service
    ports:
      - "8001:8001"
    environment:
      DB_URL: postgres://storywarz_user:secure_password@postgres:5432/storywarz_db
      REDIS_URL: redis://redis:6379

  lobby-service:
    build:
      context: ./services/lobby-service
    ports:
      - "8002:8002"
    environment:
      DB_URL: postgres://storywarz_user:secure_password@postgres:5432/storywarz_db
      REDIS_URL: redis://redis:6379

  # Additional services (matchmaking, game, scoring, prompt, notification) to be defined similarly.

  graphql-gateway:
    build:
      context: ./graphql-gateway
    ports:
      - "8080:8080"
    environment:
      USER_SERVICE_ADDR: "user-service:8001"
      LOBBY_SERVICE_ADDR: "lobby-service:8002"
      GOOGLE_OAUTH_CLIENT_ID: "your_google_oauth_client_id"
      GOOGLE_OAUTH_CLIENT_SECRET: "your_google_oauth_client_secret"
      GOOGLE_OAUTH_REDIRECT_URL: "https://your-domain.com/api/auth/google/callback"
      JWT_SECRET_KEY: "your_secure_jwt_secret_key"

  frontend:
    build:
      context: ./frontend
    ports:
      - "3000:3000"
    environment:
      API_GATEWAY_URL: "http://graphql-gateway:8080"

volumes:
  storywarz_postgres_data:
-----------------------------------------------------------

8.2 Environment Variables

- OAuth: Set GOOGLE_OAUTH_CLIENT_ID, GOOGLE_OAUTH_CLIENT_SECRET, and GOOGLE_OAUTH_REDIRECT_URL in the User Service and GraphQL Gateway.
- JWT: Use a secure JWT_SECRET_KEY for signing tokens.

===========================================================================
9. FUTURE CONSIDERATIONS

- Real-Time Updates: Implement GraphQL subscriptions for real-time notifications.
- gRPC Streaming & Interceptors: Enhance inter-service communication with streaming and interceptors.
- Schema Evolution: Use deprecation strategies and version the GraphQL API as new features are added.
- Scalability: Decouple services further and implement advanced caching techniques as the application grows.

===========================================================================
CONCLUSION

This specification provides comprehensive details for implementing the StoryWarz project:
- It covers architectural choices, data handling, error strategies, and testing plans.
- It includes guidelines for building and deploying a system using Golang, gRPC, and a GraphQL Gateway.

A developer can begin implementation immediately using this document as a guide.
===========================================================================
