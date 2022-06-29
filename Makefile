crawl:
	go run cmd/crawler/main.go


server:
	go run cmd/server/main.go

compose:
	docker compose up -d --build