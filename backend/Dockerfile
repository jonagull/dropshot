# Start with a minimal Go image
FROM golang:1.19-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go module file (only go.mod, no go.sum)
COPY go.mod ./

# Download necessary Go dependencies (if any)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Create the uploads directory (in case it doesn't exist)
RUN mkdir -p ./uploads

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
