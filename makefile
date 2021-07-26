build:
	go build -o dcli

plain: build
	./dcli

get: build	
	./dcli get

set: build	
	./dcli set

windows: 
	GOOS=windows GOARCH=amd64 go build -o dcliwin

zone: build	
	./dcli zone

lint: 
	golangci-lint run --timeout=20m

checkin: lint
	go mod tidy

install:
	sudo cp ./dcli* /usr/local/bin
	