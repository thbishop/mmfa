root_dir = `pwd`

all: fmt test
	@mkdir -p bin/
	@env bash --norc -i ./scripts/build.sh

fmt:
	@echo
	@echo "==> Formatting source code."
	@echo
	@env go fmt ./...

help:
	@echo "default\t\ttest, format, and build the code"
	@echo "fmt\t\tformat the code"
	@echo "package\t\tbuild/package the code for platforms"
	@echo "test\t\ttest the code"

package:
	@echo "\n==> Packaging the code\n"
	@mkdir -p pkg/
	@env bash --norc -i ./scripts/dist.sh

test:
	@env godep go test ./...

.PNONY: all fmt help package test
