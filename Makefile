.PHONY: go-run-cli

go-run-cli:
	@echo "Checking Docker engine status..."
	@docker info >/dev/null 2>&1 || (echo "Docker engine is not running. Please start Docker and try again." && exit 1)
	@echo "Docker engine is running."
	@echo "Starting Docker containers..."
	@docker-compose up -d
	@echo "Running Go CLI..."
	@cd cmd && go run main.go
