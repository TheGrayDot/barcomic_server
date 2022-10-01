BINARY_NAME=barcomic_server

run:
	sudo go run cmd/barcomic_server/main.go  

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/barcomic_server/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/barcomic_server/main.go

clean:
	go clean
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows

format:
	@gofmt -l .

test:
	@go test -v internal/server

test_coverage:
	@go test -cover internal/server/
