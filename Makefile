run:
	go run main.go

build:
	go build -o runServer

test:
	go test ./api/handler/login_handler_test.go

test-v:
	go test ./api/handler/login_handler_test.go -v