#!/bin/bash


cd "$(dirname "$0")" || exit 1
source ./export.sh

mkdir -p ../bin

GOARCH=amd64 \
GOOS=linux \
go build \
-ldflags \
"-X main.Version=$PROJECT_VERSION \
-X main.Hash=$COMMIT_HASH" \
-o "../bin/$BINARY_PREFIX-linux" \
../cmd/barcomic/main.go

chmod u+x "../bin/$BINARY_PREFIX-linux"
