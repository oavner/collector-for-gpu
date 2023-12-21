# Use the official golang image as a base image
FROM golang:1.19-alpine

# Set the working directory to the app directory
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o collector-for-gpu .

# Command to run the application
ENTRYPOINT ["./collector-for-gpu"]