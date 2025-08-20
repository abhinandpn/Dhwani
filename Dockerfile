# Use the latest stable Go version
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

# Copy binary and static files (no secrets!)
COPY --from=builder /app/dhwani .
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./dhwani"]
