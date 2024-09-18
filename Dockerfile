FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main usersServer/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client

COPY --from=builder /app/main main

EXPOSE 50051

CMD ["./main"]
