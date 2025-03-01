FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY .env .env

EXPOSE 5000

CMD ["/app/main"]