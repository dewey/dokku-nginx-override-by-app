GO_REPO_ROOT := /go/src/github.com/dokku/smoke-test-plugin
BUILD_IMAGE := golang:1.7.1

.PHONY: build-in-docker build clean src-clean

build-in-docker:
	docker run --rm \
		-v $$PWD:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT) \
		$(BUILD_IMAGE) \
		bash -c "make build" || exit $$?

build: commands triggers
triggers: nginx-app-template-source
commands: **/**/commands.go
	go build -a -o commands ./src/commands/commands.go

nginx-app-template-source: **/**/nginx-app-template-source.go
	go build -a -o nginx-app-template-source ./src/triggers/nginx-app-template-source.go

clean:
	rm -f commands nginx-app-template-source

src-clean:
	rm -rf .editorconfig .gitignore src LICENSE Makefile README.md *.go
