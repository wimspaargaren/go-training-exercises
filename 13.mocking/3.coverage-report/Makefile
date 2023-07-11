.PHONY: test

test:
	@go test -coverprofile=coverage.cov ./... -cover -race -p=4 -v
	@go tool cover -func=coverage.cov > functioncoverage.out
	@go tool cover -html=coverage.cov -o coverage.html
	@open ./coverage.html
