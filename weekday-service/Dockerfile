FROM golang:1.13 AS builder

RUN mkdir /app
ADD weekday.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS runner

COPY --from=builder /app .
CMD ["./main"]
