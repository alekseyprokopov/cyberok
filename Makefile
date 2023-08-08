run:
	go build ./cmd/cyberok/main.go; ./main

up:
	docker-compose up --build cyberok
