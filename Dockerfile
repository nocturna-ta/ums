# Use a multi-stage build to keep the final image size smaller
FROM golang:1.23.6-bookworm AS builder

# Install git and other required packages
RUN apt-get update && apt-get install -y git gcc build-essential

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Set up GitHub authentication for private repositories
ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# Set GOPRIVATE for private repositories
ENV GOPRIVATE=github.com/nocturna-ta/*

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ums .

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ums .

# Create a minimal image for running the application
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

# Copy the binary from the builder stage
COPY --from=builder /app/ums /app/
COPY --from=builder /app/config /app/config

# Create the config directory in case it doesn't exist
RUN mkdir -p /app/config/files

# Expose the service port
EXPOSE 8900
EXPOSE 35000

# Run the binary
ENTRYPOINT ["/app/ums"]
CMD ["server-http"]