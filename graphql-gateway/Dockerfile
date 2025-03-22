FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o graphql-gateway ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/graphql-gateway .

EXPOSE 8080

CMD ["./graphql-gateway"] 