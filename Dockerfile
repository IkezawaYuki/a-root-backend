FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o aroot cmd/server/main.go
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

FROM alpine:latest
USER appuser
WORKDIR /app
COPY --from=builder /app/aroot .
EXPOSE 8080
CMD ["./aroot"]