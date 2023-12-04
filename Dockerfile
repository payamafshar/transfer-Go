# Use official golang image from docker hub with golang version 1.20.5 as dev environment
FROM golang:1.21.4 as dev

# Set environment variables for goproxy and gopath


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

# Expose port 5050
EXPOSE 5050

CMD reflex -g '*.go' go run main.go --start-service


