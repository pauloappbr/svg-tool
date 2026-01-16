BINARY_NAME=svg-tool
MAIN_PATH=cmd/svg-tool/main.go

# Detect Operating System
OS := $(shell uname -s)

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe
	rm -rf bin/

test:
	go test ./pkg/... -v

install:
	@echo "Detected OS: $(OS)"
	@if [ "$(OS)" = "Windows_NT" ]; then \
		powershell -ExecutionPolicy Bypass -File scripts/install.ps1; \
	else \
		bash scripts/install.sh; \
	fi

# Cross-Compilation for release
release:
	@echo "Generating release binaries..."
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/$(BINARY_NAME) $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/$(BINARY_NAME).exe $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o bin/mac-arm64/$(BINARY_NAME) $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o bin/mac-amd64/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Binaries created in bin/ directory."