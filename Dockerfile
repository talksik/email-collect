FROM golang:1.21 AS builder
COPY / /app

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest AS production
WORKDIR /
COPY --from=builder /app/cmd .
CMD ["./main"]

