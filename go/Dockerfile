FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/ .

RUN go build -o voting-system .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/voting-system .

EXPOSE 3000

CMD ["./voting-system"]
