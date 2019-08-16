version = `cat VERSION`

.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: build
build:
	rm -rf ./bin && mkdir ./bin

	GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-linux-386 ./cmd/tmpl
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-linux-amd64 ./cmd/tmpl

	GOOS=darwin GOARCH=386 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-darwin-386 ./cmd/tmpl
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-darwin-amd64 ./cmd/tmpl

	GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-windows-386 ./cmd/tmpl
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-windows-amd64 ./cmd/tmpl

	GOOS=openbsd GOARCH=386 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-openbsd-386 ./cmd/tmpl
	GOOS=openbsd GOARCH=amd64 go build -ldflags "-X main.version=$(version)" -o ./bin/tmpl-openbsd-amd64 ./cmd/tmpl

.PHONY: help
help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
