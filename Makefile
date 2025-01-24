# DEFINITIONS

# BUILDING TARGETS
.PHONY: all build

all: build

build:
	@echo "Building application..."


# TESTING TARGETS
.PHONY: lint

lint:
	golangci-lint run ./...