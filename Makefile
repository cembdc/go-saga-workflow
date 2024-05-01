.PHONY: build run

build:
	@echo "Building..."
	@go build -o app cmd/main.go

run: build
	@echo "Running..."
	@./app