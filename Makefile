APP := app
MAIN := main.go
CMD_DIR := ./cmd/api
BIN_DIR := bin

PKGS := ./...
TEST_FLAGS ?= -v -race # own flags: make test TEST_FLAGS="-count=1 -race"
COVER_FILE := coverage.out

TARGETS := dev clean test coverage lint build help debug-cpu start

.PHONY: $(TARGETS)

.DEFAULT:
	@echo "Unknown target: $@"
	@echo ""
	@$(MAKE) help

help:
	@echo "Available targets:"
	@echo "  make start       - starts the app"
	@echo "  make debug-cpu   - make cpu profile debug"
	@echo "  make test        - run unit tests"
	@echo "  make coverage    - run tests with coverage report"
	@echo "  make lint        - format and lint code"
	@echo "  make build       - build binaries"
	@echo "  make clean       - remove generated files"

start:
	docker compose up -d
	go run $(CMD_DIR)/$(MAIN)

debug-cpu:
	docker compose up -d
	go run -tags debug $(PKGS)
	#go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

clean:
	rm -f $(COVER_FILE)
	rm -rf $(BIN_DIR)

test:
	go test $(TEST_FLAGS) $(PKGS)

coverage: clean test
	go test -coverpkg=$(PKGS) -coverprofile=$(COVER_FILE) $(PKGS)
	go tool cover -html=$(COVER_FILE)

lint:
	go fmt $(PKGS)
	go vet $(PKGS)
	golangci-lint run

build:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP)-linux-amd64 $(CMD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP)-windows-amd64.exe $(CMD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(APP)-darwin-arm64 $(CMD_DIR)

dev:
	@echo "Cleaning MODs..."
	go mod tidy
	@echo "Running linters..."
	$(MAKE) lint
	@echo "Running tests..."
	$(MAKE) test
	@echo "Building binaries..."
	$(MAKE) build
	@echo "Done"
