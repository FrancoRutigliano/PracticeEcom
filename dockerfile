# Use an official Golang runtime as the base image
FROM golang:1.22

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o app ./cmd/main.go

# Set the entrypoint for the container
ENTRYPOINT ["./app"]
