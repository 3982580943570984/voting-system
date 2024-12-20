FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY services/votes/go.mod services/votes/go.sum ./votes/

COPY services/shared ./shared/

WORKDIR /app/votes

RUN go mod download

WORKDIR /app

COPY services/votes ./votes/

WORKDIR /app/votes

RUN go build -ldflags="-s -w" -o votes .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/votes ./

EXPOSE 3001

CMD ["./votes"]
