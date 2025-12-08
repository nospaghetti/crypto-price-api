FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/crypto-api ./cmd/crypto-api/main.go

FROM alpine:3.22 AS crypto-api

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/bin/crypto-api .

EXPOSE 8080
CMD ["./crypto-api"]
