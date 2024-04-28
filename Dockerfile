# Start from a Node.js base image to have npm available
FROM node:20 AS node_base

# Set the current working directory
WORKDIR /app

# Copy the files
COPY . .

# Install dependencies
RUN npm install

# Start a new stage from golang base image
FROM golang:1.22

# Copy Node.js dependencies from the previous stage
COPY --from=node_base /app/node_modules /app/node_modules

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the files
COPY . .
RUN ls -la

# Enable CGO
ENV CGO_ENABLED=1

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./main"]
