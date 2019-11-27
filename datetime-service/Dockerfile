FROM golang:1.13 AS builder

RUN mkdir /app
ADD datetime.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production

COPY --from=builder /app .
CMD ["./main"]
