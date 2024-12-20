FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY . .

RUN go build -o votes .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/votes .

EXPOSE 3000

CMD ["./votes"]
