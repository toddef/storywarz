#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}StoryWarz Setup Test${NC}"
echo "Testing if Docker and Docker Compose are installed..."

# Check Docker installation
if command -v docker &> /dev/null; then
    echo -e "${GREEN}[✓] Docker is installed${NC}"
    docker --version
else
    echo -e "${RED}[✗] Docker is not installed${NC}"
    echo "Please install Docker before continuing."
    exit 1
fi

# Check Docker Compose installation
if command -v docker-compose &> /dev/null; then
    echo -e "${GREEN}[✓] Docker Compose is installed${NC}"
    docker-compose --version
else
    echo -e "${RED}[✗] Docker Compose is not installed${NC}"
    echo "Please install Docker Compose before continuing."
    exit 1
fi

echo ""
echo "Testing if required ports are available..."

# Test ports
declare -a ports=(5432 6379 8080 50051)
for port in "${ports[@]}"; do
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null ; then
        echo -e "${RED}[✗] Port $port is already in use${NC}"
    else
        echo -e "${GREEN}[✓] Port $port is available${NC}"
    fi
done

echo ""
echo "Starting containers in detached mode..."
docker-compose up -d postgres redis

# Wait for services to be ready
echo "Waiting for PostgreSQL to be ready..."
sleep 5

# Test PostgreSQL connection
echo ""
echo "Testing connection to PostgreSQL..."
if docker exec storywarz-postgres pg_isready -U storywarz > /dev/null; then
    echo -e "${GREEN}[✓] Successfully connected to PostgreSQL${NC}"
else
    echo -e "${RED}[✗] Could not connect to PostgreSQL${NC}"
fi

# Test Redis connection
echo ""
echo "Testing connection to Redis..."
if docker exec storywarz-redis redis-cli -a storywarz_password ping | grep -q "PONG"; then
    echo -e "${GREEN}[✓] Successfully connected to Redis${NC}"
else
    echo -e "${RED}[✗] Could not connect to Redis${NC}"
fi

echo ""
echo "Stopping containers..."
docker-compose down

echo ""
echo -e "${YELLOW}Setup test completed!${NC}"
echo "You can start all services with: docker-compose up"
echo "To run in the background, use: docker-compose up -d" 