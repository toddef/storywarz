.PHONY: proto build test run clean

# Default target
all: proto build

# Generate protobuf code
proto:
	@echo "Generating protocol buffer code..."
	@mkdir -p pkg/proto
	@protoc --go_out=. --go-grpc_out=. proto/user.proto
	@echo "Done."

# Build the application
build:
	@echo "Building user-service..."
	@go build -o bin/user-service ./cmd/server
	@echo "Done."

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./test/...
	@echo "Done."

# Run the application
run:
	@echo "Running user-service..."
	@go run ./cmd/server/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f bin/user-service
	@echo "Done." 