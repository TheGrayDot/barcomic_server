#!/bin/bash


source scripts/export.sh

GOARCH=amd64 \
GOOS=linux \
go build -o bin/$BINARY_PREFIX-linux internal/server

chmod u+x bin/$BINARY_PREFIX-linux
