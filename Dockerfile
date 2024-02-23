# Use the official Golang image as the base image
FROM golang:1.22.0-bookworm

# Set the working directory inside the container
WORKDIR /goGinApp

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Copy .env file
COPY .env .env

# Build the Go application
RUN go build -o main .

# Expose the port that your Go application will run on
EXPOSE 9090

# Command to run your Go application
CMD ["./main"]
