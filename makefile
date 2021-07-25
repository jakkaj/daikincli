build:
	go build -o dcli

run: build	
	./dcli get

lint: 
	golangci-lint run --timeout=20m

checkin: lint
	go mod tidy