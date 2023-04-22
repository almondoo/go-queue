start:
	docker compose build --no-cache

build:
	docker compose build

up:
	docker compose up -d

restart:
	docker compose restart

down:
	docker compose down

ps:
	docker compose ps

go:
	docker compose exec go-queue ash

godoc:
	docker compose exec go-queue godoc -http ":8081"

redis:
	docker compose exec redis ash

db:
	docker compose exec db mysql -u root --password=root

l-go:
	docker compose logs go-queue -f
	