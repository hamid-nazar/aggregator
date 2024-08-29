PROJECT = ../rss-aggregator

# Run the Go file without creating a binary
run:
	@echo "Running $(MAIN_FILE)..."
	@go run $(PROJECT)
	
# Run the Go file and create a binary
lint:
	@echo "Linting the project..."
	@golangci-lint run

