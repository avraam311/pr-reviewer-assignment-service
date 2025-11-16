.PHONY: lint, up, buildup, down, test-integration

lint:
	go vet ./...
	golangci-lint run ./...

up:
	docker compose up

buildup:
	docker compose up --build

down:
	docker compose down -v

test-integration:
	go run test_integration.go
