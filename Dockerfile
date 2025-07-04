# Stage 1 - Build (Creation of binary)
FROM golang:1.24 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

# Copying source code
COPY src src

# Building application
RUN go build -o ./japa_tracker src/main.go

# Stage 2 - Execution 
FROM alpine:latest

# Copy binary
COPY --from=build /app/japa_tracker .

ENTRYPOINT ["./japa_tracker"]
