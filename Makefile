SHELL := /bin/bash
BINARY=echo-template
export PATH := $(CURDIR)/_tools/bin:$(PATH)
export GOFLAGS := -mod=vendor
USER ?= `whoami`

.PHONY: all
.DEFAULT: all

all: mod test

tidy:
	go mod tidy

mod:
	go mod vendor

docs:
	# will create docs

test:
	go test


format: mod
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run -v --timeout 5m --out-format checkstyle > golangci-report.xml
.PHONY: lint


install:
	# linter
	go build -o _tools/bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	# swagger
	go build -o _tools/bin/swag github.com/swaggo/swag/cmd/swag

local: mod
	go run cmd/template/main.go

docker.run: docker.echo

docker.echo:
	docker build -t echo-template . && docker run -p 8080:8080 echo-template

$(BINARY):
	go build ${LDFLAGS} -o ${BINARY} ./cmd/template

