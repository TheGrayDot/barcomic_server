#!/bin/bash


source scripts/export.sh

go clean
rm -f "bin/$BINARY_PREFIX-linux"
rm -f "bin/$BINARY_PREFIX-windows.exe"
rm -f "bin/$BINARY_PREFIX-darwin"
