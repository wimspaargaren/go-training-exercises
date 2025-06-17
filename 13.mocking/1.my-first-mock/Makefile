.PHONY: mocks

$(GOBIN)/mockery:
	@go install github.com/vektra/mockery/v3@v3.4.0
	@go mod tidy

mocks: | $(GOBIN)/mockery
	@mockery
