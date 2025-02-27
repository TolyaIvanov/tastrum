В миграциях заполнил немного некоторые таблицы, т.к. они завязаны на создании и применении промокодов, четко можно
увидеть как в миграциях, так и в файле promo_repository

Сборка
docker compose build --no-cache
docker compose up -d

Поднять, опустить миграции
docker compose exec app sh -c 'migrate -path /app/migrations -database "$DB_URL" up'
docker compose exec app sh -c 'migrate -path /app/migrations -database "$DB_URL" down'

все апи
GET http://localhost:8083/check-health работает ли апи
GET http://localhost:8083/api/promocode/:code Применение промокода
GET http://localhost:8083/api/rewards возвращает весь список наград
GET http://localhost:8083/api/GetPlayers возвращает весь список наград

POST http://localhost:8083/admin/promocode создаем новый промокод

GET http://localhost:8083/admin/ веб страница


Открывается сайт
http://localhost:8083/admin/