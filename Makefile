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

# Deploy ECR stack
deploy-ecr-cf:
	aws cloudformation deploy \
		--template-file cloudformation/ecr-cloudformation.yaml \
		--stack-name oauth-service-ecr \
		--region eu-west-1 \
		--capabilities CAPABILITY_NAMED_IAM \
		--no-fail-on-empty-changeset

# Deploy application stack
deploy-apprunner-cf:
	aws cloudformation deploy \
		--template-file cloudformation/apprunner-cloudformation.yaml \
		--stack-name oauth-service-application \
		--region eu-west-1 \
		--capabilities CAPABILITY_NAMED_IAM \
		--no-fail-on-empty-changeset

# Deploy both stacks
deploy-all: deploy-ecr-cf deploy-apprunner-cf

# Delete application stack
delete-application:
	aws cloudformation delete-stack \
		--stack-name oauth-service-application \
		--region eu-west-1

# Delete ECR stack
delete-ecr-cf:
	aws cloudformation delete-stack \
		--stack-name oauth-service-ecr \
		--region eu-west-1

# Delete both stacks
delete-all: delete-application delete-ecr-cf

# Check if the service is healthy
health:
	@echo "Checking OAuth service health..."
	@curl -f http://localhost:8080/health || (echo "❌ Service is not healthy" && exit 1)
	@echo "✅ Service is healthy"
