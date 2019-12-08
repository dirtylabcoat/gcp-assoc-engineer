FROM golang:1.13 AS builder

RUN mkdir /app
ADD quote.go /app/
WORKDIR /app
RUN go get -u github.com/go-sql-driver/mysql
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS runner

COPY --from=builder /app .
CMD ["./main"]
