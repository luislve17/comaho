run:
	docker compose up --build -d


down:
	docker compose down

restart: down
	docker compose up --build -d
