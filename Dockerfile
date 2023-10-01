FROM golang AS builder
COPY / /app

WORKDIR /app/cmd
RUN CGO_ENABLED=1 GOOS=linux go build -o main

FROM alpine:latest AS production
COPY --from=builder /app/cmd .
CMD ["./main"]
