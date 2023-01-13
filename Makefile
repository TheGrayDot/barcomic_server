run:
	@go run cmd/main.go -v

clean:
	./scripts/build_clean.sh

install_golang_deps:
	@go get ./internal/barcomic

install_linux_deps:
	./scripts/install_linux_deps.sh

update_golang_packages:
	./scripts/update_golang_packages.sh

format:
	@gofmt -l .

test:
	@go test -v ./internal/barcomic

coverage:
	@go test -cover ./internal/barcomic

build_linux:
	./scripts/build_linux.sh

build_windows:
	./scripts/build_windows.sh

build_darwin:
	./scripts/build_darwin.sh
