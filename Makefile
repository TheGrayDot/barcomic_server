run:
	@go run cmd/main.go

clean:
	./scripts/build_clean.sh

install_golang_deps:
	@go get internal/barcomic_server

install_linux_deps:
	./scripts/install_linux_deps.sh

format:
	@gofmt -l .

test:
	@go test -v internal/barcomic_server

coverage:
	@go test -cover internal/barcomic_server

build_linux:
	./scripts/build_linux.sh

build_windows:
	./scripts/build_windows.sh
