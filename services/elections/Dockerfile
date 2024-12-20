FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY services/elections/go.mod services/elections/go.sum ./elections/

COPY services/shared ./shared/

WORKDIR /app/elections

RUN go mod download

WORKDIR /app

COPY services/elections ./elections/

WORKDIR /app/elections

RUN go build -ldflags="-s -w" -o elections .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/elections ./

EXPOSE 3002

CMD ["./elections"]
