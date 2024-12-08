# Stage 1: Build Backend
FROM golang:1.20 as backend-builder
WORKDIR /app

ARG SERVICE
COPY ${SERVICE}/ /app/

RUN if [ "$SERVICE" = "backend" ]; then go mod download && go build -o main .; fi

# Stage 2: NGINX for Load Balancing
FROM nginx:alpine as nginx-builder
WORKDIR /etc/nginx
COPY nginx/nginx.conf /etc/nginx/nginx.conf

# Final Stage: Use appropriate runtime
FROM alpine:latest
WORKDIR /app

# Backend runtime
COPY --from=backend-builder /app/main .
COPY --from=backend-builder /app/go.mod .
COPY --from=backend-builder /app/go.sum .
EXPOSE 8080
CMD ["./main"]

# NGINX runtime
COPY --from=nginx-builder /etc/nginx /etc/nginx
CMD ["nginx", "-g", "daemon off;"]
