# Simple Makefile for the Go web server application

.PHONY: build run clean test dev docker-build docker-run docker-stop docker-clean docker-logs

# Build the application
build:
	go build -o ./bin/buildconfig-server .

# Run the web server
run:
	go run .

# Run in development mode (with auto-reload)
dev:
	@echo "Starting web server in development mode..."
	@echo "Server will be available at: http://localhost:8080"
	go run .

# Clean build artifacts
clean:
	rm -f buildconfig-server

# Run tests (if any)
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy

# Test the web server endpoints
test-endpoints:
	@echo "Testing web server endpoints..."
	@echo "1. Testing home page..."
	@curl -s http://localhost:8080/ | head -5
	@echo "\n2. Testing /api/time..."
	@curl -s http://localhost:8080/api/time
	@echo "\n3. Testing /api/health..."
	@curl -s http://localhost:8080/api/health
	@echo "\n4. Testing /api/unzip..."
	@curl -s http://localhost:8080/api/unzip

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t buildconfig-server .

docker-run:
	@echo "Running Docker container..."
	docker run -d --name buildconfig-server -p 8080:8080 buildconfig-server

docker-stop:
	@echo "Stopping Docker container..."
	docker stop buildconfig-server || true
	docker rm buildconfig-server || true

docker-clean:
	@echo "Cleaning Docker resources..."
	docker stop buildconfig-server || true
	docker rm buildconfig-server || true
	docker rmi buildconfig-server || true

docker-logs:
	@echo "Showing Docker container logs..."
	docker logs -f buildconfig-server


# Test Docker container endpoints
test-docker-endpoints:
	@echo "Testing Docker container endpoints..."
	@echo "1. Testing home page..."
	@curl -s http://localhost:8080/ | head -5
	@echo "\n2. Testing /api/time..."
	@curl -s http://localhost:8080/api/time
	@echo "\n3. Testing /api/health..."
	@curl -s http://localhost:8080/api/health
	@echo "\n4. Testing /api/unzip..."
	@curl -s http://localhost:8080/api/unzip
