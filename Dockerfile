# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install standard build dependencies if needed (e.g., git, make)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
# CGO_ENABLED=0 for static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for external API calls if needed
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/app/infra/database/migrations ./app/infra/database/migrations
COPY --from=builder /app/config ./config

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main", "app"]
