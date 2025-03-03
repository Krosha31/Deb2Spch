# Этап сборки
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы приложения
COPY . .

# Компилируем приложение для Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/app/main.go

# Этап запуска
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/myapp .

# Копируем директорию web в контейнер
COPY --from=builder /app/web /web
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/addons /addons

# Указываем команду для запуска приложения
CMD ["./myapp"]
