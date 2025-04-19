FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bootstrap ./cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bootstrap .

EXPOSE $APP_PORT

CMD ["./bootstrap"]
