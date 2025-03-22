FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o matchmaking-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/matchmaking-service .

EXPOSE 50053

CMD ["./matchmaking-service"] 