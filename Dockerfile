FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o qa-service cmd/app/main.go

FROM debian:latest
WORKDIR /root/
COPY --from=builder /app/qa-service .
CMD ["./qa-service"]