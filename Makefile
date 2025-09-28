default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Build the provider
.PHONY: build
build:
	go build -o terraform-provider-vapi

# Install the provider locally
.PHONY: install
install: build
	mkdir -p ~/.terraform.d/plugins/local/faruk/vapi/1.0.0/darwin_amd64/
	cp terraform-provider-vapi ~/.terraform.d/plugins/local/faruk/vapi/1.0.0/darwin_amd64/

# Clean build artifacts
.PHONY: clean
clean:
	rm -f terraform-provider-vapi

# Run tests
.PHONY: test
test:
	go test ./... -v

# Format code
.PHONY: fmt
fmt:
	go fmt ./...
	terraform fmt -recursive ./examples/

# Run linter
.PHONY: lint
lint:
	golangci-lint run

# Generate documentation
.PHONY: docs
docs:
	go generate ./...

# Initialize go modules
.PHONY: mod
mod:
	go mod tidy
	go mod download

# Run all checks (format, lint, test)
.PHONY: check
check: fmt lint test

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build     - Build the provider binary"
	@echo "  install   - Install the provider locally"
	@echo "  test      - Run unit tests"
	@echo "  testacc   - Run acceptance tests"
	@echo "  fmt       - Format code"
	@echo "  lint      - Run linter"
	@echo "  docs      - Generate documentation"
	@echo "  mod       - Tidy and download go modules"
	@echo "  check     - Run format, lint, and test"
	@echo "  clean     - Remove build artifacts"
	@echo "  help      - Show this help message"
