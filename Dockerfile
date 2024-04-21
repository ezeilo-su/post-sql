FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./src ./src

RUN go mod download
RUN go mod verify

RUN go build -ldflags="-w -s" -o server ./src/server

FROM scratch
COPY --from=builder app/server bin/

ENTRYPOINT [ "/bin/server" ]