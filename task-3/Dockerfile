FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#Runner
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]
