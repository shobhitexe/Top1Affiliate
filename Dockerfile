
FROM golang:1.24-alpine AS builder


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN GOMAXPROCS=1 go build -o main ./cmd/server


FROM alpine:latest


WORKDIR /root/


COPY --from=builder /app/main .


EXPOSE 8080


CMD ["./main"]
