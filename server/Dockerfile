FROM golang:1.24-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/api

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY configs/prod.yaml /app/configs/prod.yaml

ENV CONFIG_PATH=/app/config/dev.yaml

EXPOSE 3000

CMD ["./app"]
