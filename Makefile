.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD..."
	@go mod vendor
	@go mod tidy

run:
	@go run cmd\main.go