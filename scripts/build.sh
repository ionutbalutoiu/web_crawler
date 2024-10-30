#!/usr/bin/env bash
set -e

DIR="$(dirname $0)"

BUILD_DIR="${DIR}/../build"
CMD_DIR="${DIR}/../cmd"

mkdir -p $BUILD_DIR

echo "Building web_crawler"
go build -o ${BUILD_DIR}/web_crawler -ldflags="-extldflags '-static' -s -w" ${CMD_DIR}/web_crawler

BIN_PATH="$(realpath ${BUILD_DIR}/web_crawler)"
echo "Binary available at: ${BIN_PATH}"
