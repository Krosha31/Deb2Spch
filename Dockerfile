FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/myapp .

COPY --from=builder /app/web /web
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/addons /addons

CMD ["./myapp"]
