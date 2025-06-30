.PHONY: lint

$(GOBIN)/golangci-lint:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.0

lint: | $(GOBIN)/golangci-lint
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=1 run

$(GOBIN)/gofumpt:
	@GO111MODULE=on go get mvdan.cc/gofumpt
	@go mod tidy

gofumpt: | $(GOBIN)/gofumpt
	@gofumpt -w .
	@gofumpt -w ./mypkg
