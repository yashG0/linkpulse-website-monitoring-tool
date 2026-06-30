FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o linkpulse ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/linkpulse .

COPY .env .

EXPOSE 8080

CMD ["./linkpulse"]