build:
	go build -o dcli

run: build	
	./dcli get

zone: build	
	./dcli zone

lint: 
	golangci-lint run --timeout=20m

checkin: lint
	go mod tidy