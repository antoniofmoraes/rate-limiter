FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/rate-limiter

FROM scratch
COPY --from=builder /app/server .
COPY --from=builder /app/cmd/rate-limiter/*.env .
CMD ["./server"]