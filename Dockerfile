# Use the golang image as a base
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o app ./main.go

# Start with a fresh, minimal image
FROM gcr.io/distroless/base-debian10

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage to the container
COPY --from=builder /app/app .

# Set the command to run the binary
CMD ["/app/app"]



