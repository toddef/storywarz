Step 1: Repository Setup & Infrastructure
Goal:
Set up the repository structure and basic Docker Compose configuration. This will create the foundation for our microservices and environment.

Tasks:

Create the folder structure for the monorepo (services for user, lobby, matchmaking, game, scoring, prompt, notification, and the GraphQL gateway).
Create placeholder files for each service (Dockerfile, README, proto folder, etc.).
Create a top-level docker-compose.yaml to start PostgreSQL, Redis, and service placeholders.
Write initial integration tests (e.g., test that the containers can start).
Prompt 1: Repository Structure & Docker Compose

text
Copy
Create a repository structure for the StoryWarz project as follows:

/storywarz
 ├── docker-compose.yaml
 ├── spec.md
 ├── services/
 │    ├── user-service/
 │    │     ├── proto/user.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    ├── lobby-service/
 │    │     ├── proto/lobby.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    ├── matchmaking-service/
 │    │     ├── proto/matchmaking.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    ├── game-service/
 │    │     ├── proto/game.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    ├── scoring-service/
 │    │     ├── proto/scoring.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    ├── prompt-service/
 │    │     ├── proto/prompt.proto
 │    │     ├── Dockerfile
 │    │     └── README.md
 │    └── notification-service/
 │          ├── proto/notification.proto
 │          ├── Dockerfile
 │          └── README.md
 └── graphql-gateway/
        ├── schema.graphql
        ├── resolvers.go
        ├── gqlgen.yml
        ├── Dockerfile
        └── README.md

Additionally, create a `docker-compose.yaml` that starts:
- PostgreSQL (with initial environment variables and volume for persistence)
- Redis
- Placeholder containers for at least the user-service and graphql-gateway

Include basic instructions and initial tests (for example, a script or instructions that confirm containers can start and connect).

The output should be a complete repository scaffolding with the above structure and a working docker-compose file.
Step 2: Base Golang Service – User Service
Goal:
Develop the basic Golang structure for the user-service. This service will later expose gRPC endpoints.

Tasks:

Create a basic main package for the user-service.
Set up a gRPC server that listens on a designated port (e.g., 8001).
Write an initial unit test that ensures the server starts without error.
Prepare a placeholder proto file (user.proto) with a minimal service definition.
Prompt 2: User Service Base Setup

text
Copy
For the user-service, set up a basic Golang project that does the following:
1. Implements a main package that starts a gRPC server listening on port 8001.
2. Provides a placeholder proto file (proto/user.proto) with a basic service definition for UserService (e.g., including a GetUser RPC).
3. Write a simple unit test that verifies the gRPC server starts correctly and is reachable.

Ensure the code follows best practices (modular, error handling, logging) and is ready for further expansion.
Step 3: Implement User Service gRPC Method
Goal:
Implement the first gRPC method in the user-service. We’ll create a GetUser method that returns a static user record for testing.

Tasks:

Extend the proto file to define a GetUser request and response.
Generate the Go code using protoc with Go plugins.
Implement the GetUser method in the user-service server that returns a static user.
Write unit tests to call GetUser via a gRPC client and assert the returned data is as expected.
Prompt 3: User Service GetUser Implementation

text
Copy
Extend the user-service project as follows:
1. In proto/user.proto, define the GetUser RPC, along with GetUserRequest and GetUserResponse messages.
2. Generate the Go code using protoc (with protoc-gen-go and protoc-gen-go-grpc).
3. Implement the GetUser method in the user-service gRPC server so that it returns a static user record (with fields like id, email, name, and profilePicture).
4. Write unit tests that create a gRPC client, call GetUser, and validate that the correct static data is returned.

Ensure that error handling and context timeouts are properly set.
Step 4: Set Up GraphQL Gateway with gqlgen
Goal:
Initialize the GraphQL Gateway project using gqlgen with a minimal schema. Create a basic "me" query that returns a User object.

Tasks:

Set up a new Golang project under graphql-gateway.
Install and configure gqlgen with an initial schema file (schema.graphql).
Define a minimal schema that includes a User type and a query named me.
Generate resolver stubs using gqlgen.
Implement a dummy resolver for the me query that returns a static user.
Write tests for the GraphQL endpoint to verify the me query works.
Prompt 4: GraphQL Gateway Base Setup

text
Copy
For the GraphQL Gateway:
1. Set up a new Golang project in the graphql-gateway folder.
2. Configure gqlgen (create a gqlgen.yml and schema.graphql) with a minimal schema that defines:
   - A User type with id, email, name, and profilePicture fields.
   - A Query type with a "me" field that returns a User.
3. Run gqlgen to generate resolver stubs.
4. Implement a dummy resolver for the "me" query that returns a static user.
5. Write unit tests (using a testing framework or gqlgen’s playground) to confirm that querying "me" returns the expected static user.

Ensure that the implementation follows a schema-first approach and is test-driven.
Step 5: Integrate gRPC Client in GraphQL Gateway
Goal:
Wire up the GraphQL Gateway to communicate with the user-service via gRPC. Update the resolver for the me query to call the GetUser method from user-service.

Tasks:

In the GraphQL Gateway, set up a gRPC client for the user-service.
Update the me resolver to:
Extract the user ID from the request context (simulate authentication).
Call the user-service’s GetUser method via gRPC.
Map the returned data to the GraphQL User type.
Write integration tests for the me query that verify the end-to-end flow (GraphQL query → gRPC call → static user returned).
Prompt 5: GraphQL Gateway gRPC Integration

text
Copy
Enhance the GraphQL Gateway as follows:
1. In the GraphQL Gateway project, add code to initialize a gRPC client that connects to the user-service (using a configurable address from environment variables).
2. Update the resolver for the "me" query so that instead of returning a static user, it:
   a. Retrieves a user ID from the request context (simulate a valid authenticated user).
   b. Calls the GetUser RPC on the user-service via gRPC.
   c. Maps the response from the gRPC call to the GraphQL User type.
3. Write integration tests that perform a GraphQL query on "me" and validate that the response correctly reflects the data from the user-service.

Ensure error handling and timeouts are implemented for the gRPC call.
Step 6: Iterative Service Expansion (e.g., Lobby Service)
Goal:
After the user-service and GraphQL gateway integration are working, expand the architecture to add another service (for example, the lobby-service). The process should mirror the steps taken for the user-service.

Tasks:

Set up the basic scaffolding for the lobby-service (similar to Step 2).
Define the proto file for lobby-service with RPCs (e.g., CreateLobby, JoinLobby, GetLobby).
Implement one RPC method (e.g., CreateLobby) with static response.
Add unit tests for the lobby-service.
In the GraphQL Gateway, extend the schema with queries/mutations for lobby operations.
Wire up the lobby-service gRPC client in the GraphQL resolvers.
Write integration tests for lobby-related GraphQL operations.
Prompt 6: Lobby Service Implementation and Integration

text
Copy
Repeat the process for the lobby-service:
1. Create the base Golang project for lobby-service in its designated folder.
2. Define proto/lobby.proto with RPCs like CreateLobby, JoinLobby, and GetLobby.
3. Generate the Go gRPC code and implement at least the CreateLobby RPC to return a static lobby with a join code.
4. Write unit tests for the CreateLobby method.
5. Update the GraphQL Gateway:
   a. Extend schema.graphql to include mutations/queries for lobby operations (e.g., createLobby, joinLobby).
   b. Add code to initialize a gRPC client for the lobby-service.
   c. Implement the corresponding resolvers that call the lobby-service via gRPC.
6. Write integration tests to validate that GraphQL mutations for creating or joining a lobby work as expected.

Keep the steps small and ensure each part is test-driven before moving on.
Step 7: Testing, Integration, and Continuous Wiring
Goal:
Ensure all components are integrated and fully tested. Set up integration tests and a CI pipeline to run tests on each commit.

Tasks:

Write integration tests that start all services via Docker Compose and verify end-to-end flows (e.g., user authentication via GraphQL, lobby creation, etc.).
Set up a CI pipeline (or at least document the steps) to run unit tests, integration tests, and code linting.
Perform manual testing using tools like Postman or GraphQL Playground to verify complete flows.
Ensure that every service is wired correctly in the Docker Compose file.
Prompt 7: Integration Testing & CI Setup

text
Copy
Develop a testing and integration plan for the complete project:
1. Write integration tests that:
   a. Spin up the full stack (using Docker Compose) including PostgreSQL, Redis, user-service, lobby-service, and the GraphQL Gateway.
   b. Execute end-to-end scenarios (e.g., query the "me" endpoint, create a lobby, join a lobby) and verify responses.
2. Document steps for setting up a CI pipeline that runs all tests (unit and integration) on every commit.
3. Ensure that each test case cleans up after itself and does not leave any orphaned data.
4. Provide instructions on how to run tests locally and via CI.

The output should be a detailed testing plan along with example test cases (code snippets) that verify each service's integration.
Step 8: Final Wiring, Cleanup & Documentation
Goal:
Finalize the project by ensuring all components are integrated, documented, and ready for further development or deployment.

Tasks:

Review all services to ensure they are connected via the GraphQL Gateway.
Add documentation in each service’s README.md detailing the purpose, API endpoints, and how to run tests.
Update the top-level spec.md file with any final notes.
Do a final round of manual testing and integration testing.
Prompt 8: Final Integration & Documentation

text
Copy
Perform final integration and cleanup tasks:
1. Review and ensure that all services (user, lobby, matchmaking, game, scoring, prompt, notification) are correctly wired to the GraphQL Gateway via gRPC.
2. Update each service’s README.md with:
   - A description of the service.
   - Instructions for building and running the service.
   - API documentation (including details from proto files).
   - Testing instructions.
3. Finalize the top-level spec.md file with all architectural decisions, integration details, and testing plans.
4. Write a final integration test that runs through a full user journey (e.g., authentication, lobby creation, and joining) via the GraphQL Gateway.
5. Document any next steps or potential future enhancements.

The final output should be a fully integrated, documented, and test-driven codebase that can be built and deployed using Docker Compose.
