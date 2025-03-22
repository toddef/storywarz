FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o game-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/game-service .

EXPOSE 50054

CMD ["./game-service"] 