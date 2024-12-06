deps:
	go mod download

lint:
	gofmt -w -s .

templates:
	templ generate

test:
	templ generate & go test -cover ./...

build:
	templ generate & go build -o ./tmp/main cmd/main.go

start:
	templ generate & go run cmd/main.go

dev:
	air