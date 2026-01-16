# Stage 1: Builder
FROM golang:1.23-alpine AS builder

# Set environment variables for static build
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copy dependency files first (caching layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -ldflags="-s -w" -o svg-tool cmd/svg-tool/main.go

# Stage 2: Final Image
FROM alpine:latest

# Install basic certificates (good practice)
RUN apk --no-cache add ca-certificates

WORKDIR /workdir

# Copy binary from builder
COPY --from=builder /app/svg-tool /usr/local/bin/svg-tool

# Define Entrypoint
ENTRYPOINT ["svg-tool"]

# Default argument if none provided
CMD ["-help"]