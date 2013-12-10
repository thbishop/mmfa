DEPS = $(go list -f '{ range .TestImports}}{ .}} { end}}' ./...' } } }')

all: fmt test
	@mkdir -p bin/
	@bash --norc -i ./scripts/build.sh

deps:
	go get -d -v ./...
	echo $(DEPS) | xargs -n1 go get -d

fmt:
	@echo "\n==> Formatting source code"
	@go fmt ./...

help:
	@echo "default\t\ttest, format, and build the code"
	@echo "deps\t\tget dependencies"
	@echo "fmt\t\tformat the code"
	@echo "package\t\tbuild/package the code for platforms"
	@echo "test\t\ttest the code"

package:
	@echo "\n==> Packaging the code\n"
	@mkdir -p pkg/
	@bash --norc -i ./scripts/dist.sh

test:
	go list ./... | xargs -n1 go test

.PNONY: all fmt help package test
