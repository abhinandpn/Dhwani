# Use the latest stable Go version that supports >=1.23.6
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o dhwani ./main.go

# Final lightweight image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/dhwani .
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./dhwani"]
# Copy env and credentials
COPY server/dhwani-key.json /app/server/dhwani-key.json
COPY .env /app/.env
