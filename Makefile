down:
	docker compose down

restart: down
	docker compose up --build comaho

test:
	docker compose up --build test


