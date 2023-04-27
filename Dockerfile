# Use a golang base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the entire contents of the current directory into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the container
EXPOSE 8080

# Run the Go application
CMD ["./main"]