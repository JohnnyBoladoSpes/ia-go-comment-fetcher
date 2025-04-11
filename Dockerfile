FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o comment-fetcher .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/comment-fetcher .
COPY assets/ ./assets/

EXPOSE 8080
CMD ["./comment-fetcher"]