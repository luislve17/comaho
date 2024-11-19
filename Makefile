down:
	docker compose down

restart: down
	docker compose up --build -d comaho

test:
	docker compose up --build test


