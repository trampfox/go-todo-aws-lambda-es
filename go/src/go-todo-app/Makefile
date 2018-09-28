clean: 
	rm -rf ./bin/get
	rm -rf ./bin/list
	rm -rf ./bin/post
	rm -rf ./bin/auth

build-local:
	GOOS=darwin go build -o bin/get ./functions/get.go
	GOOS=darwin go build -o bin/list ./functions/list.go
	GOOS=darwin go build -o bin/post ./functions/post.go
	GOOS=darwin go build -o bin/auth ./functions/auth.go
	chmod +x bin/*

build:
	GOOS=linux GOARCH=amd64 go build -o bin/get ./functions/get.go
	GOOS=linux GOARCH=amd64 go build -o bin/list ./functions/list.go
	GOOS=linux GOARCH=amd64 go build -o bin/post ./functions/post.go
	GOOS=linux GOARCH=amd64 go build -o bin/auth ./functions/auth.go
	chmod +x bin/*
