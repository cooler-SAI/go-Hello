FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main -ldflags "-s -w" .
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]