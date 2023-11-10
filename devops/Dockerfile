FROM golang:1.21.3-alpine3.18 AS builder

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/myapp ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build/myapp .
COPY --from=builder /app/internal/config/config.yml .
CMD ["./myapp"]