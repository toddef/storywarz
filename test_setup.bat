@echo off
setlocal

echo StoryWarz Setup Test
echo Testing if Docker and Docker Compose are installed...

:: Check Docker installation
docker --version >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    echo [v] Docker is installed
    docker --version
) else (
    echo [x] Docker is not installed
    echo Please install Docker before continuing.
    exit /b 1
)

:: Check Docker Compose installation
docker-compose --version >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    echo [v] Docker Compose is installed
    docker-compose --version
) else (
    echo [x] Docker Compose is not installed
    echo Please install Docker Compose before continuing.
    exit /b 1
)

echo.
echo Testing if required ports are available...
:: Note: This is simplified as Windows doesn't have an easy equivalent to lsof
echo [!] Port availability check not implemented on Windows.
echo [!] Please ensure ports 5432, 6379, 8080, and 50051 are available.

echo.
echo Starting containers in detached mode...
docker-compose up -d postgres redis

:: Wait for services to be ready
echo Waiting for PostgreSQL to be ready...
timeout /t 5 /nobreak >nul

:: Test PostgreSQL connection
echo.
echo Testing connection to PostgreSQL...
docker exec storywarz-postgres pg_isready -U storywarz >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    echo [v] Successfully connected to PostgreSQL
) else (
    echo [x] Could not connect to PostgreSQL
)

:: Test Redis connection
echo.
echo Testing connection to Redis...
docker exec storywarz-redis redis-cli -a storywarz_password ping >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    echo [v] Successfully connected to Redis
) else (
    echo [x] Could not connect to Redis
)

echo.
echo Stopping containers...
docker-compose down

echo.
echo Setup test completed!
echo You can start all services with: docker-compose up
echo To run in the background, use: docker-compose up -d

endlocal 