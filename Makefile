down:
	docker compose down

restart: down
	docker compose up --build -d
