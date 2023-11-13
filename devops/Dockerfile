FROM golang:1.21.3-alpine3.18 AS builder

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/simple-stock ./cmd

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/build/simple-stock .
COPY --from=builder /app/internal/config/config.yml /app/internal/config/
CMD ["./simple-stock"]