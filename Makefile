.PHONY: build
build:
	docker-compose build --no-cache

.PHONY: up
up:
	docker-compose up

.PHONY: upd
upd:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: exec
exec:
	docker-compose exec ngo bash

.PHONY: gobuild
gobuild:
	cp config.yml bin/
	GOOS=linux go build -ldflags="-s -w" -o ./bin/ndeio cmd/nde-clock-io/main.go

.PHONY: m-gobuild
m-gobuild:
	go build -o ./bin/ndeio cmd/nde-clock-io/main.go

.PHONY: run-in
run-in:
	go run cmd/nde-clock-io/main.go clockin

.PHONY: run-out
run-out:
	go run cmd/nde-clock-io/main.go clockout
