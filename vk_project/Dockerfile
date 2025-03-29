FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o voting-bot cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/voting-bot .

EXPOSE 8080

CMD ["./voting-bot"]
