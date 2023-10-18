#!/bin/bash


source scripts/export.sh

mkdir bin

GOARCH=amd64 \
GOOS=linux \
go build \
-ldflags \
"-X main.Version=$PROJECT_VERSION \
-X main.Hash=$COMMIT_HASH" \
-o "bin/$BINARY_PREFIX-linux" \
cmd/barcomic/main.go

chmod u+x "bin/$BINARY_PREFIX-linux"
