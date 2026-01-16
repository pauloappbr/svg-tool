BINARY_NAME=svg-tool
MAIN_PATH=cmd/svg-tool/main.go

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

install: build
	@echo "Detected OS: $(OS)"
	@if [ "$(OS)" = "Windows_NT" ]; then \
		powershell -ExecutionPolicy Bypass -File scripts/install.ps1; \
	else \
		sudo bash scripts/install.sh; \
	fi

release:
	@echo "Generating release binaries..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/linux-amd64/$(BINARY_NAME) $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/windows-amd64/$(BINARY_NAME).exe $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/mac-arm64/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Binaries created in bin/ directory."