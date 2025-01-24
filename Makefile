# COMMON DEFINITIONS
## COLORS
NoColor=\033[0m
White=\033[0;37m
Green=\033[0;32m
Yellow=\033[0;33m
Red=\033[0;31m

## PROJECT SPECIFIED VARIABLES
BuildDir=bin


# BUILDING TARGETS
.PHONY: all build build-echo

all: build

build: build-echo

build-echo:
	@printf "${Green}[echo]${NoColor} Building...\n"
	GOOS=linux GOARCH=amd64 go build -o ./${BuildDir}/echo-amd64 ./cmd/echo/*.go
	GOOS=windows GOARCH=amd64 go build -o ./${BuildDir}/echo-amd64.exe ./cmd/echo/*.go


# TESTING TARGETS
.PHONY: lint

lint:
	golangci-lint run ./...