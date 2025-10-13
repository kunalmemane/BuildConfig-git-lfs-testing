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
   ./buildconfig-server
   ```

4. **Development mode:**
   ```bash
   make dev
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

- `make run` - Run the web server
- `make dev` - Run in development mode
- `make build` - Build the application
- `make clean` - Clean build artifacts
- `make test` - Run tests
- `make test-endpoints` - Test all API endpoints
- `make deps` - Install dependencies

## Project Structure

- `main.go` - Web server entry point and route definitions
- `utils.go` - HTTP handlers and utility functions
- `go.mod` - Go module definition
- `Makefile` - Build automation and testing

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
