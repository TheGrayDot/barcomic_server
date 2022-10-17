BINARY_NAME=barcomic_server

run:
	@go run cmd/barcomic_server/main.go  

clean:
	go clean
	rm -f bin/${BINARY_NAME}-linux
	rm -f bin/${BINARY_NAME}-windows

format:
	@gofmt -l .

test:
	@go test -v internal/server

coverage:
	@go test -cover internal/server/

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/barcomic_server/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/barcomic_server/main.go
