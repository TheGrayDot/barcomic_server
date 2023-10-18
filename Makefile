run:
	@go run cmd/barcomic/main.go -v

clean:
	./scripts/build_clean.sh

docker_run:
	docker image build -t barcomic .; \
	docker run --name barcomic -p 8080:80 barcomic

docker_clean:
	docker container stop barcomic; \
	docker container rm barcomic; \
	docker image rm barcomic

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
