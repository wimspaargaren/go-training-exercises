.PHONY: mocks

$(GOBIN)/mockery:
	@go install github.com/vektra/mockery/v2@latest
	@go mod tidy

mocks: | $(GOBIN)/mockery
	@mockery --output . --filename data_store_mock_test.go --inpackage --name DataStore
