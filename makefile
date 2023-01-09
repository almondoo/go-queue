start:
	docker compose build --no-cache

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

ps:
	docker compose ps

go:
	docker compose exec go-queue ash

db:
	docker compose exec db mysql -u root --password=root