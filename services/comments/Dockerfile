FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY services/comments/go.mod services/comments/go.sum ./comments/

COPY services/shared ./shared/

WORKDIR /app/comments

RUN go mod download

WORKDIR /app

COPY services/comments ./comments/

WORKDIR /app/comments

RUN go build -ldflags="-s -w" -o comments .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/comments ./

EXPOSE 3001

CMD ["./comments"]
