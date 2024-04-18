############################
# STEP 1 build executable binary
############################
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY ./src ./src
RUN go mod download
RUN go mod verify

# Build the binary.
RUN go build -ldflags="-w -s" -o bin ./src/server/

############################
# STEP 2 build a small image
############################
FROM alpine
# FROM scratch

# Copy our static executable
COPY --from=builder bin .

# Run the hello binary.
ENTRYPOINT ["/bin"]
