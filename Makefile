.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD TIDY..."
	@go mod tidy

	@echo "RUNNING GO MOD VENDOR..."
	@go mod vendor

run:
	@go run cmd\main.go