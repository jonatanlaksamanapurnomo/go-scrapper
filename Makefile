.PHONY: go-run-cli

go-run-cli:
	docker-compose up -d
	cd cmd && go run main.go
