.PHONE:

build:
	go build -o ./.bin/cmd cmd/main.go

run: build
	./.bin/cmd

