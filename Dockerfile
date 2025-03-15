# Build
FROM golang:1.24-bullseye AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

# Run
FROM debian:bullseye-20250224-slim
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]