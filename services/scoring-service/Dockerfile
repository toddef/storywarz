FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o scoring-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/scoring-service .

EXPOSE 50055

CMD ["./scoring-service"] 