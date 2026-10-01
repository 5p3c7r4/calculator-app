# Multi-stage Dockerfile for calculator app
# Build frontend and backend separately

# Stage 1: Build frontend
FROM node:20-alpine AS frontend-build

WORKDIR /app/frontend
COPY calculator-frontend/package*.json ./
RUN npm ci

COPY calculator-frontend/ ./

RUN npm run build

# Stage 2: Build backend
FROM golang:1.26.2-alpine AS backend-build

WORKDIR /app/backend
COPY calculator-api/go.mod calculator-api/go.sum ./
RUN go mod download

COPY calculator-api/ ./

RUN go build -o calculator-api cmd/api/main.go

# Stage 3: Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy built frontend
COPY --from=frontend-build /app/frontend/dist ./frontend
# Copy built backend
COPY --from=backend-build /app/backend/calculator-api ./calculator-api

# Expose ports
EXPOSE 8000

# Create a simple script to start both services
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["/root/entrypoint.sh"]