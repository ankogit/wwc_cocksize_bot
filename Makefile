.PHONY:

build:
	go build -o ./.bin/bot cmd/main.go

run: build
	./.bin/bot

build-image:
	docker build -t wwc_cocksize_bot:0.1 .

start-container:
	docker run --env-file .env -p 80:80 wwc_cocksize_bot:0.1

build-compose:
	docker-compose

run-compose:
	docker-compose up

build-run-compose:
	docker-compose up -d --build

gen:
	protoc --proto_path=proto --go_out=api --go_opt=paths=source_relative --go-grpc_out=. proto/*.proto