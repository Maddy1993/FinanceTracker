# Dockerfile

# ---------- Build Stage ----------
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build your gRPC server binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/base
COPY --from=builder /app/server /server

# Default gRPC port
EXPOSE 50051

# Run the server
CMD ["/server"]
