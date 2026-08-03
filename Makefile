.PHONY: lb-demo cache-demo rate-demo down

lb-demo:
	cd src/01-load-balancing-demo && docker compose up --build

cache-demo:
	cd src/02-cache-aside-demo && docker compose up --build

rate-demo:
	cd src/03-rate-limiting-demo && docker compose up --build

down:
	cd src/01-load-balancing-demo && docker compose down -v
	cd src/02-cache-aside-demo && docker compose down -v
	cd src/03-rate-limiting-demo && docker compose down -v
