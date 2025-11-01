# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies for CGO (required for SQLite)
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build both binaries
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/api ./cmd/http
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/cli ./cmd/cli

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite-libs libc6-compat

WORKDIR /app

COPY --from=builder /app/bin/api /app/bin/api
COPY --from=builder /app/bin/cli /app/bin/cli

COPY config.json /app/config.json
COPY db /app/db

RUN mkdir -p /app/data

EXPOSE 8080

CMD ["/app/bin/api"]
