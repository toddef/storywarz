FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o prompt-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/prompt-service .

EXPOSE 50056

CMD ["./prompt-service"] 