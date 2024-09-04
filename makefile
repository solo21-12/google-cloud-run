test:
	go test -v ./Tests/...
down:
	docker compose down 
up:
	docker compose up -d --build
logs:
	docker exec -it gcr-api tail -f /var/log/gcr-api.err.log