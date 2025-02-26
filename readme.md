docker compose build --no-cache
docker compose up -d

docker compose exec app sh -c 'migrate -path /app/migrations -database "$DB_URL" up'
docker compose exec app sh -c 'migrate -path /app/migrations -database "$DB_URL" down'
