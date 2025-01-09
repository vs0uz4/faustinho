FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cli ./cmd

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/ .

COPY .env .

EXPOSE 8080

CMD ["./cli"]
