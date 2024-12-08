# Real-Time Collaborative Editor

A scalable, distributed real-time collaborative text editor with multi-user support, conflict resolution, and version control, similar to Google Docs.

## Technologies Used

- **Backend**: Golang, WebSockets, Redis
- **Frontend**: HTML, CSS, JavaScript
- **Load Balancer**: NGINX
- **Hosting**: AWS S3 (for frontend)
- **Containerization**: Docker, Docker Compose

## Features

- Real-time collaborative editing with WebSocket synchronization.
- Conflict resolution using Operational Transformation (OT).
- Multi-user support with live cursor tracking.
- Redis for shared state management.
- NGINX for load balancing backend services.
- Frontend hosted on AWS S3.

## Prerequisites

- **Docker** and **Docker Compose** installed.
- An AWS S3 bucket for hosting the frontend.

## Environment Variables

Create a `.env` file in the project root with the following content:

```plaintext
# Frontend Configuration
AWS_S3_URL=https://your-bucket-name.s3.amazonaws.com

# Backend Configuration
REDIS_HOST=redis:6379
SERVER_PORT=8080

# NGINX Configuration
BACKEND_API_URL=http://backend1:8080/ws
