#!/bin/bash


source scripts/export.sh

GOARCH=amd64 \
GOOS=linux \
go build \
-ldflags \
"-X internal/barcomic_server.Version=$PROJECT_VERSION \
-X internal/barcomic_server.Hash=$COMMIT_HASH" \
-o "bin/$BINARY_PREFIX-linux" \
cmd/main.go

chmod u+x "bin/$BINARY_PREFIX-linux"
