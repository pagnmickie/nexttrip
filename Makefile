GO_PKGS=$(shell go list ./... | grep -v -e "/scripts")
GO_PROXY_BUILD_DATE_TIME=$(shell date -u "+%Y.%m.%d %H:%M:%S %Z")
GO_PROXY_VERSION ?= UNSET
GO_PROXY_BRANCH ?= UNSET
GO_PROXY_COMMIT ?= UNSET

format: check-gofmt test

build: go-build

go-build:
	@CGO_ENABLED=0 go build -i -ldflags='-X "git.target.com/api-platform/go-proxy/proxy.version=$(GO_PROXY_VERSION)" -X "git.target.com/api-platform/go-proxy/proxy.buildDateTime=$(GO_PROXY_BUILD_DATE_TIME)" -X "git.target.com/api-platform/go-proxy/proxy.branch=$(GO_PROXY_BRANCH)" -X "git.target.com/api-platform/go-proxy/proxy.revision=$(GO_PROXY_COMMIT)"'

check-gofmt:
	@echo "Checking formatting..."
	@FMT="0"; \
	for pkg in $(GO_PKGS); do \
		OUTPUT=`gofmt -l $(GOPATH)/src/$$pkg/*.go`; \
		if [ -n "$$OUTPUT" ]; then \
			echo "$$OUTPUT"; \
			FMT="1"; \
		fi; \
	done ; \
	if [ "$$FMT" -eq "1" ]; then \
		echo "Problem with formatting in files above."; \
		exit 1; \
	else \
		echo "Success - way to run gofmt!"; \
	fi

test:
	@go test -coverprofile=c.out $(GO_PKGS) && go tool cover -html=c.out -o coverage.html

functional:
	@go test -functional git.target.com/api-platform/go-proxy
