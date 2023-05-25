mkfile_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
version := $(git tag -l | tail -n 1)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

debug: ## Build a debug version
	@mkdir -p $(mkfile_path)build
	@go build main.go -o $(mkfile_path)open-go-knocking
	@echo "Debug release done"

release: clean-build build-linux build-darwin build-windows ## Build the release version

clean-build: ## Clean the build directory
	@rm -fr @mkdir -p $(mkfile_path)build/*

build-linux: ## Build release version for GNU/Linux
	@mkdir -p $(mkfile_path)build
	@GOOS="linux" GOARCH="amd64" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-linux-amd64 main.go
	@GOOS="linux" GOARCH="arm64" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-linux-arm64 main.go
	@GOOS="linux" GOARCH="386" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-linux-386 main.go

build-darwin: ## Build release version for MacOS
	@mkdir -p $(mkfile_path)build
	@GOOS="darwin" GOARCH="amd64" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-darwin-amd64 main.go
	@GOOS="darwin" GOARCH="arm64" go build -ldflags="-X 'main.version=v0.0.1'" -o $(mkfile_path)open-go-knocking-$(version)-darwin-arm64 main.go

build-windows: ## Build release version for Windows
	@mkdir -p $(mkfile_path)build
	@GOOS="windows" GOARCH="amd64" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-windows-amd64.exe main.go
	@GOOS="windows" GOARCH="386" go build -ldflags="-X 'main.version=$(version)'" -o $(mkfile_path)open-go-knocking-$(version)-windows-386.exe main.go
