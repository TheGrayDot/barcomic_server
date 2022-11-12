#!/bin/bash


source scripts/export.sh

GOARCH=amd64 \
GOOS=linux \
go build \
-ldflags \
"-X internal/server.Version=$PROJECT_VERSION \
-X internal/server.Hash=$COMMIT_HASH" \
-o bin/$BINARY_PREFIX-linux \
cmd/barcomic_server/main.go

chmod u+x bin/$BINARY_PREFIX-linux
