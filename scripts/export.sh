#!/bin/bash


BINARY_PREFIX="barcomic_server"
PROJECT_VERSION="v0.1.0"
COMMIT_HASH=$(git rev-parse HEAD)

export BINARY_PREFIX
export PROJECT_VERSION
export COMMIT_HASH
