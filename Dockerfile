FROM golang:1.24-alpine AS builder

# Set working dir
WORKDIR /app

# Back to root
WORKDIR /app

# Copy Go mod and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the app directories
COPY cmd/ ./cmd/
COPY config ./config/
COPY internal ./internal/

# Build final binary
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/api /app/cmd/api/main.go

# Final stage — smaller image
FROM debian:bookworm-slim

RUN addgroup --system nonroot \
    && adduser --system --ingroup nonroot nonroot

WORKDIR /app

COPY --from=builder /bin/api /bin/api
COPY etc/config/server.dev.env ./etc/config/server.dev.env

USER nonroot

EXPOSE 8080
