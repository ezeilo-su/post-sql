FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY ./ ./

RUN go build -ldflags="-w -s" -o server ./server

FROM scratch
COPY --from=builder app/server /usr/server

ENTRYPOINT [ "/usr/server" ]
