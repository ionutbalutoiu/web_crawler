#!/usr/bin/env bash
set -e

DIR="$(dirname $0)"

BUILD_DIR="${DIR}/../build"
CMD_DIR="${DIR}/../cmd"

mkdir -p $BUILD_DIR

BIN_PATH="$(realpath ${BUILD_DIR}/web_crawler)"

echo "Building web_crawler at: ${BIN_PATH}"
go build -o ${BIN_PATH} -ldflags="-extldflags '-static' -s -w" ${CMD_DIR}/web_crawler
