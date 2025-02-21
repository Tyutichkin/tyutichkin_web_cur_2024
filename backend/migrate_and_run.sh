#!/bin/sh
set -e

# Ожидаем, пока Postgres станет доступным
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "db" -U "postgres" -d "mydb" -c '\q'; do
  echo "Postgres ещё не готов - ждем..."
  sleep 2
done

echo "Postgres готов, запускаем миграции..."

goose -dir /migrations postgres "$DATABASE_URL" up

echo "Миграции применены. Запускаем приложение..."
exec /app
