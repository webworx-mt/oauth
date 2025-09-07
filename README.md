# OAuth Service

A Go-based OAuth service application deployed on AWS using App Runner, ECR, and CloudFormation.

## 🏗️ Architecture

This service is built with:
- **Backend**: Go 1.25.1
- **Deployment**: AWS App Runner
- **Container Registry**: Amazon ECR
- **Infrastructure**: AWS CloudFormation
- **CI/CD**: GitHub Actions

## 📁 Project Structure

```
oauth/
├── cloudformation/           # AWS CloudFormation templates
│   ├── ecr-cloudformation.yaml      # ECR repository and IAM roles
│   └── apprunner-cloudformation.yaml # AppRunner service configuration
├── .github/workflows/        # GitHub Actions workflows
│   ├── deploy-ecr-cf.yml           # Deploy ECR infrastructure
│   ├── deploy-ecr.yml              # Build and push Docker images
│   └── deploy-apprunner-cf.yml     # Deploy AppRunner service
├── handlers/                 # HTTP request handlers
│   ├── health.go             # Health check endpoint
│   └── user.go               # User management endpoints
├── models/                   # Data models
│   ├── health.go             # Health response model
│   └── user.go               # User model
├── router/                   # HTTP routing
│   └── router.go             # Route definitions
├── server/                   # Server configuration
│   └── server.go             # Server startup logic
├── main.go                   # Application entry point
├── Dockerfile                # Container build instructions
├── docker-compose.yaml       # Local development setup
├── go.mod                    # Go module dependencies
└── Makefile                  # Build and deployment commands
```

## 🚀 API Endpoints

### Health Check
- **GET** `/health` - Returns service health status

### User Management
- **GET** `/users` - Returns list of users
- **GET** `/users/{id}` - Returns specific user by ID

## 🛠️ Local Development

### Prerequisites
- Go 1.25.1+
- Docker & Docker Compose

### Running Locally

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd oauth
   ```

2. **Install dependencies**
   ```bash
   make tidy
   ```

3. **Run with Docker Compose**
   ```bash
   make run
   ```
   The service will be available at `http://localhost:8080`

4. **Or build and run manually**
   ```bash
   make build
   ./oauth-service
   ```

### Available Make Commands

```bash
make run              # Run with Docker Compose
make build            # Build Docker image
make clean            # Clean Docker artifacts
make tidy             # Install Go dependencies
make health           # Check service health
```

## ☁️ AWS Deployment

### Infrastructure Components

1. **ECR Repository** (`oauth-service-ecr`)
   - Container image registry
   - Lifecycle policy (keeps last 10 images)
   - IAM role for AppRunner access

2. **AppRunner Service** (`oauth-service-application`)
   - Auto-scaling container service
   - Public endpoint
   - Health check configuration

### Deployment Workflows

The project uses three GitHub Actions workflows:

1. **`deploy-ecr-cf.yml`** - Deploys ECR infrastructure
   - Triggers: Changes to `cloudformation/ecr-cloudformation.yaml`
   - Creates ECR repository and IAM roles

2. **`deploy-ecr.yml`** - Builds and pushes Docker images
   - Triggers: After ECR CloudFormation deployment
   - Builds Go application and pushes to ECR

3. **`deploy-apprunner-cf.yml`** - Deploys AppRunner service
   - Triggers: After ECR image build
   - Deploys AppRunner service using latest image

### Manual Deployment

You can also deploy manually using the Makefile:

```bash
# Deploy ECR infrastructure
make deploy-ecr-cf

# Deploy AppRunner service
make deploy-apprunner-cf

# Deploy both stacks
make deploy-all
```

### Required AWS Secrets

Configure these secrets in your GitHub repository:

- `AWS_ACCESS_KEY_ID` - AWS access key
- `AWS_SECRET_ACCESS_KEY` - AWS secret key

## 🔧 Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)

### AWS Configuration

- **Region**: `eu-west-1`
- **ECR Repository**: `oauth-service`
- **AppRunner Service**: `oauth-service`

## 📊 Health Monitoring

The service provides a health check endpoint at `/health` that returns:

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "service": "oauth-service"
}
```

## 🧪 Testing

Check if the service is running:

```bash
make health
```

Or manually:

```bash
curl http://localhost:8080/health
```

## 📝 Development Notes

- The service uses mock data for user endpoints
- Health check endpoint is used by AppRunner for service health monitoring
- Docker images are tagged with both commit SHA and version tags
- CloudFormation stacks are deployed with IAM capabilities for role creation

## 🔄 CI/CD Pipeline

The deployment pipeline follows this sequence:

1. **ECR Infrastructure** → Creates ECR repository and IAM roles
2. **Image Build** → Builds and pushes Docker image to ECR
3. **AppRunner Deployment** → Deploys service using the new image

Each step triggers the next, creating a complete automated deployment pipeline.