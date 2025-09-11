.PHONY: run build clean tidy health

# Run the application using Docker
run: tidy
	docker-compose up --build

# Build the Docker image
build: tidy
	docker build -t oauth-service .

# Clean Docker artifacts
clean:
	docker-compose down
	docker rmi oauth-service 2>/dev/null || true

# Install dependencies
tidy:
	go mod tidy

# Check if the service is healthy
health:
	@echo "Checking OAuth service health..."
	@curl -f http://localhost:8080/health || (echo "❌ Service is not healthy" && exit 1)
	@echo "✅ Service is healthy"
