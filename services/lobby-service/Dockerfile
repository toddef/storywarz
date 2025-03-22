FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o lobby-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/lobby-service .

EXPOSE 50052

CMD ["./lobby-service"] 