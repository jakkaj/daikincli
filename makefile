build:
	go build -o dcli

lint: 
	golangci-lint run --timeout=20m

checkin: lint
	go mod tidy