FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# FROM postgres:latest

# ENV POSTGRES_DB=postgres
# ENV POSTGRES_USER=root
# ENV POSTGRES_PASSWORD=ChangeMe

# COPY migrations/sql/init_database.sql /docker-entrypoint-initdb.d/

COPY --from=builder /app/myapp .

# EXPOSE 5432
# EXPOSE 8080

# CMD ["sh", "-c", "docker-entrypoint.sh postgres & myapp"]
CMD ["./myapp"]

