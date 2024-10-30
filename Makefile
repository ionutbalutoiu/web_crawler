default: build

# Builds binary native to the host OS and architecture
.PHONY: build
build:
	./scripts/build.sh

# Builds binary for Linux x86_64
.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 ./scripts/build.sh

# Builds binary for MacOS ARM64
.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 ./scripts/build.sh

# Run unit tests
.PHONY: test
test:
	go test -v -race -timeout=15m -parallel=4 -count=1 ./...
