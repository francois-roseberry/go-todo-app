deps:
	go mod download

lint:
	gofmt -w -s .

templ:
	templ generate

test:
	templ generate & go test -v -cover ./...

build:
	templ generate & go build -o ./tmp/main cmd/main.go

start:
	templ generate & go run cmd/main.go

dev:
	air