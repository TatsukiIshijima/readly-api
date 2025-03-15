FROM golang:1.24-bullseye
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

EXPOSE 8080
CMD ["/app/main"]