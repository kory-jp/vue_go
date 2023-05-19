up:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

build:
	docker compose build

ps:
	docker compose ps

login:
	docker exec -it go_container /bin/sh