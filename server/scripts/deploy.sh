#!/bin/bash

# Production deployment script
set -e

# Configuration
PROJECT_NAME="myapp"
DOCKER_IMAGE="$PROJECT_NAME:latest"
COMPOSE_FILE="docker-compose.prod.yml"

echo "Starting deployment for $PROJECT_NAME"

# Check if .env.production exists
if [ ! -f .env.production ]; then
    echo "Error: .env.production file not found"
    exit 1
fi

# Load production environment
export $(cat .env.production | grep -v '#' | xargs)

# Build Docker image
echo "Building Docker image..."
docker build -f docker/Dockerfile -t "$DOCKER_IMAGE" .

# Backup database
echo "Creating database backup..."
docker-compose -f "$COMPOSE_FILE" exec -T db pg_dump -U postgres "$DB_NAME" | gzip > "backup_pre_deploy_$(date +%Y%m%d_%H%M%S).sql.gz"

# Pull latest images
echo "Pulling latest images..."
docker-compose -f "$COMPOSE_FILE" pull

# Deploy with zero downtime
echo "Deploying application..."
docker-compose -f "$COMPOSE_FILE" up -d --no-deps app

# Wait for health check
echo "Waiting for application to be healthy..."
sleep 10

# Check health
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "Application is healthy"
else
    echo "Application health check failed"
    echo "Rolling back..."
    docker-compose -f "$COMPOSE_FILE" rollback
    exit 1
fi

# Clean up old images
echo "Cleaning up old Docker images..."
docker image prune -f

echo "Deployment completed successfully!"