# Base image
FROM golang:1.19-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

# Download module dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /scholarship .

# Create minimal image
FROM scratch

# Copy the built binary into the image
COPY --from=0 /scholarship /scholarship

# Expose port 3000
EXPOSE 3000

# Run the app
CMD ["/scholarship"]

