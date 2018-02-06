bundle:
		dep ensure

build:
		go build -o bin/go-api-server-traning

run:
		go run ./*.go