lint:
	gofmt -w -s .

test:
	go test ./...

build:
	go build -o ./tmp/main cmd/main.go

start:
	go run cmd/main.go

dev:
	air