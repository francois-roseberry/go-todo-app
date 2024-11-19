lint:
	gofmt -w -s .

test:
	go test -v ./...

build:
	templ generate & go build -o ./tmp/main cmd/main.go

start:
	templ generate & go run cmd/main.go

dev:
	air