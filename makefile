build:
	go build -o dcli

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