FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bot ./cmd

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/bot .
COPY --from=builder /app/internal/src/images ./internal/src/images
COPY --from=builder /app/.env .

CMD ["./bot"]