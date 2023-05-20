run:
	go run main.go

build:
	go build -o runServer

test:
	for test in $(shell find . -name "*_test.go" -printf '%h\n' | sort -u); do \
		go test -v $$test;\
	done