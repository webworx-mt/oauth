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

# Deploy infrastructure stack
deploy-infrastructure:
	aws cloudformation deploy \
		--template-file cloudformation/infrastructure-cloudformation.yaml \
		--stack-name oauth-service-infrastructure \
		--region eu-west-1 \
		--capabilities CAPABILITY_NAMED_IAM \
		--no-fail-on-empty-changeset

# Deploy application stack
deploy-application:
	aws cloudformation deploy \
		--template-file cloudformation/application-cloudformation.yaml \
		--stack-name oauth-service-application \
		--region eu-west-1 \
		--capabilities CAPABILITY_NAMED_IAM \
		--no-fail-on-empty-changeset

# Deploy both stacks
deploy-all: deploy-infrastructure deploy-application

# Delete application stack
delete-application:
	aws cloudformation delete-stack \
		--stack-name oauth-service-application \
		--region eu-west-1

# Delete infrastructure stack
delete-infrastructure:
	aws cloudformation delete-stack \
		--stack-name oauth-service-infrastructure \
		--region eu-west-1

# Delete both stacks
delete-all: delete-application delete-infrastructure

# Check if the service is healthy
health:
	@echo "Checking OAuth service health..."
	@curl -f http://localhost:8080/health || (echo "❌ Service is not healthy" && exit 1)
	@echo "✅ Service is healthy"
