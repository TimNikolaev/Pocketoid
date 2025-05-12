.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run:	build
	./.bin/bot

build-image:
	docker build -t pocketoid-bot:v0.1 .

start-container:
	docker run --name pocketoid-bot -p 8888:8888 --env-file .env pocketoid-bot:v0.1
