FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bin ./cmd/server/
CMD ["./bin"]
