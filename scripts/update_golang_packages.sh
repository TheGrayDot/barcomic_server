#!/bin/bash


go get -u ./...
go mod tidy
cd internal/barcomic_server
go get -u ./...
go mod tidy
