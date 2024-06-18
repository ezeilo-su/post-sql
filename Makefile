.PHONY: dev test

# Development build
dev:
	docker-compose -f docker-compose-dev.yml build --parallel

# Test
test:
	docker-compose -f docker-compose-test.yml up -d
	go run cmd/main.go
	docker-compose -f docker-compose-test.yml down
