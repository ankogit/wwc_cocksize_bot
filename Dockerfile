# Используем базовый образ для Go
FROM golang:1.17.6-alpine

# Создадим директорию
RUN mkdir /app

# Скопируем всё в директорию
ADD . /app/

# Установим рабочей папкой директорию
WORKDIR /app


# Соберём приложение
RUN go build -o main ./cmd

# Запустим приложение
CMD ["/app/main"]