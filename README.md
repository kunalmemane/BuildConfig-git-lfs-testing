# BuildConfig-git-lfs-testing

A simple Go web server application for testing BuildConfig with Git LFS.

## Features

- **Web Server**: HTTP server running on port 8080
- **REST API**: JSON endpoints for various operations
- **Interactive Home Page**: HTML interface with links to all endpoints
- **Archive Processing**: Extract and list contents of tar.xz files
- **Health Check**: Service health monitoring endpoint

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page with interactive interface |
| GET | `/api/time` | Current time and timezone |
| GET | `/api/health` | Health check status |
| GET | `/api/unzip` | Unzip and list crc-linux-amd64.tar.xz contents |

## Getting Started

### Prerequisites

- Go 1.24.6 or later

### Running the Web Server

#### Option 1: Run Locally with Go

1. **Using Go directly:**
   ```bash
   go run .
   ```

2. **Using Make:**
   ```bash
   make run
   ```

3. **Building the application:**
   ```bash
   make build
   ./bin/buildconfig-server
   ```

4. **Development mode:**
   ```bash
   make dev
   ```

#### Option 2: Run with Docker

1. **Using Docker directly:**
   ```bash
   # Build the Docker image
   docker build -t buildconfig-server .
   
   # Run the container
   docker run -d --name buildconfig-server -p 8080:8080 buildconfig-server
   ```

2. **Using Make commands:**
   ```bash
   # Build and run with Docker
   make docker-build
   make docker-run
   
   # View logs
   make docker-logs
   
   # Stop and clean up
   make docker-stop
   make docker-clean
   ```


### Testing the API

Once the server is running, you can:

1. **Visit the home page**: http://localhost:8080
2. **Test API endpoints directly**:
   ```bash
   curl http://localhost:8080/api/time
   curl http://localhost:8080/api/health
   curl http://localhost:8080/api/unzip
   ```

3. **Use the Makefile test command**:
   ```bash
   make test-endpoints
   ```

### Available Make Commands

#### Local Development
- `make run` - Run the web server
- `make dev` - Run in development mode
- `make build` - Build the application
- `make clean` - Clean build artifacts
- `make test` - Run tests
- `make test-endpoints` - Test all API endpoints
- `make deps` - Install dependencies

#### Docker Commands
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container
- `make docker-stop` - Stop Docker container
- `make docker-clean` - Clean Docker resources
- `make docker-logs` - View Docker container logs
- `make test-docker-endpoints` - Test Docker container endpoints


## Project Structure

- `main.go` - Web server entry point and route definitions
- `utils.go` - HTTP handlers and utility functions
- `go.mod` - Go module definition
- `Makefile` - Build automation and testing
- `Dockerfile` - Docker image configuration

## Example API Responses

**GET /api/unzip**
```json
{
  "file": "crc-linux-amd64.tar.xz",
  "status": "success",
  "timestamp": 1703123456,
  "contents": [
    "crc (Size: 12345678 bytes, Type: 0)",
    "README.md (Size: 1024 bytes, Type: 0)"
  ],
  "file_count": 2,
  "total_size": 12346702,
  "archive_type": "tar.xz"
}
```
