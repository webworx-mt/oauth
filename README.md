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
├── .github/workflows/        # GitHub Actions workflows
│   └── deploy.yml            # Triggers centralized deployment
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
├── deployment.yaml           # Deployment configuration
├── Dockerfile                # Container build instructions
├── docker-compose.yaml       # Local development setup
├── go.mod                    # Go module dependencies
└── Makefile                  # Local development commands
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

This service uses a centralized deployment system that handles all deployment complexity automatically.

### Automatic Deployment

The service is deployed automatically when you push to the `main` or `master` branch. The deployment is handled by the centralized deployment system at [webworx-mt/deployment](https://github.com/webworx-mt/deployment).

### Configuration

The deployment is configured via the `deployment.yaml` file in the root directory:

```yaml
app:
  name: oauth-service
  type: go-apprunner
  version: "1.0.0"

deployment:
  environment: production
  region: eu-west-1
  port: 8080
  cpu: "1 vCPU"
  memory: "2 GB"
  scaling:
    min_instances: 1
    max_instances: 10
    target_cpu: 70
  health_check:
    path: "/health"
    interval: 30
    timeout: 5
    healthy_threshold: 2
    unhealthy_threshold: 3
```

### Required GitHub Secrets

Configure this secret in your GitHub repository:

- `DEPLOYMENT_TOKEN` - GitHub Personal Access Token for triggering deployments

### Manual Deployment

You can trigger a manual deployment:

1. Go to [Actions](https://github.com/webworx-mt/deployment/actions)
2. Run "Deploy Go App Runner Application"
3. Enter your repository name (`webworx-mt/oauth`)
4. Click "Run workflow"

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

The deployment pipeline is handled by the centralized deployment system:

1. **Trigger** → Push to main/master branch or manual trigger
2. **Centralized Deployment** → [webworx-mt/deployment](https://github.com/webworx-mt/deployment) handles:
   - ECR repository creation and lifecycle policies
   - Docker image build and push
   - AppRunner service deployment
   - Infrastructure management via CloudFormation

This creates a simple, consistent deployment process across all applications.