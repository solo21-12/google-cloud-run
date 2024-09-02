test:
	go test -v ./Tests/...
down:
	docker compose down 
up:
	docker compose up -d
logs:
	docker compose logs -f