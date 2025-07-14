FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o leaseweb-challenge ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/leaseweb-challenge .
# Copy config files or static assets if needed
EXPOSE 8080
CMD ["./leaseweb-challenge"]
