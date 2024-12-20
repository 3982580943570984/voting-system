FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY services/users/go.mod services/users/go.sum ./users/

COPY services/shared ./shared/

WORKDIR /app/users

RUN go mod download

WORKDIR /app

COPY services/users ./users/

WORKDIR /app/users

RUN go build -o users .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/users ./

EXPOSE 3001

CMD ["./users"]
