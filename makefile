build:
	go build -o bin/main main.go
run:
	go run .
test:
	go test

up:
	docker-compose up --detach
stop:
	docker-compose stop
restart:
	docker-compose stop && docker-compose up --detach