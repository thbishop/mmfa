gopath = "$(root_dir)/third_party:$(GOPATH)"

all: fmt test
	@mkdir -p bin/
	@env GOPATH=$(gopath) bash --norc -i ./scripts/build.sh

deps:
	@echo
	@echo "==> Downloading dependencies."
	@echo
	@env GOPATH=$(gopath) go get -d -v ./...
	@echo
	@echo "==> Removing .git and .bzr from third_party."
	@echo
	@find ./third_party -type d -name .git | xargs rm -rf
	@find ./third_party -type d -name .bzr | xargs rm -rf

fmt:
	@echo
	@echo "==> Formatting source code."
	@echo
	@env GOPATH=$(gopath) go fmt ./...

help:
	@echo "default\t\ttest, format, and build the code"
	@echo "deps\t\tget dependencies"
	@echo "fmt\t\tformat the code"
	@echo "package\t\tbuild/package the code for platforms"
	@echo "test\t\ttest the code"

package:
	@echo "\n==> Packaging the code\n"
	@mkdir -p pkg/
	@env GOPATH=$(gopath) bash --norc -i ./scripts/dist.sh

test:
	@env GOPATH=$(gopath) go test ./...

.PNONY: all fmt help package test
