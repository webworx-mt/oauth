.PHONY: run build clean tidy deploy-stack delete-stack health

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

# Deploy CloudFormation stack
deploy-stack:
	aws cloudformation deploy \
		--template-file cloudformation.yaml \
		--stack-name oauth-service-stack \
		--region eu-west-1 \
		--capabilities CAPABILITY_NAMED_IAM \
		--no-fail-on-empty-changeset

# Delete CloudFormation stack
delete-stack:
	aws cloudformation delete-stack \
		--stack-name oauth-service-stack \
		--region eu-west-1

# Check if the service is healthy
health:
	@echo "Checking OAuth service health..."
	@curl -f http://localhost:8080/health || (echo "❌ Service is not healthy" && exit 1)
	@echo "✅ Service is healthy"
