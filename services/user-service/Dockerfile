FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o user-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 50051

CMD ["./user-service"] 