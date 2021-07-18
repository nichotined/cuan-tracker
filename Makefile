SHELL=/bin/bash
APP=cuan-tracker
APP_VERSION=$(shell cat .version)
APP_COMMIT=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./build/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /build | cut -d "/" -f4 | uniq | awk '{print}' ORS='/... ')
DB_HOST=$(shell cat test.yml | grep -i DB_HOST | cut -d " " -f2)
DB_NAME=$(shell cat test.yml | grep -i DB_NAME | cut -d " " -f2)
DB_USER=$(shell cat test.yml | grep -i DB_USER | cut -d " " -f2)
DB_PORT=$(shell cat test.yml | grep -i DB_PORT | cut -d " " -f2)
COVERAGE_MIN=10
.PHONY: build

all: clean test

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

clean:
	@echo "> cleaning up the mess"
	@rm -rf build && mkdir -p build

fmt:
	@gofmt -l -s $(SOURCE_DIRS) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

build:
	@echo "> building binary"
	@go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)" cmd/*.go
	
test: generate
	@echo "> running test and creating coverage report"
	@ENVIRONMENT=test \
	go test -race -p=1 -cover -coverprofile=coverage.out $(ALL_PACKAGES)
	@go tool cover -html=coverage.out -o build/coverage.html
	@go tool cover -func=coverage.out | grep -i total:
	@go tool cover -func=coverage.out | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)"; exit 1}}'
	@rm *.out

setup:
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0

generate:
	@echo "Running go generate..."
	@go generate ./...
