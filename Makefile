run:
	@go run cmd/barcomic/main.go -v

clean:
	./scripts/build_clean.sh

install_golang_deps:
	@go get ./internal/barcomic

install_linux_deps:
	./scripts/install_linux_deps.sh

update_golang_packages:
	@go get -u ./...; go mod tidy

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

release:
	./scripts/git_tag_and_release.sh
