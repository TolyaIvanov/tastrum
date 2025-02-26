FROM golang:1.24-alpine AS builder

# Установка зависимостей для сборки
RUN apk add --no-cache git postgresql-client \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache postgresql-client curl

# Копирование бинарника и миграций
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/server /app/server
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/.env /app/.env
COPY migrations /migrations

CMD ["/app/server"]
