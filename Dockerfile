FROM golang:1.25-alpine as Builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o crypto-api .

FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/crypto-api .

EXPOSE 80

CMD ["./crypto-api"]
