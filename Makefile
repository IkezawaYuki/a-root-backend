.PHONY: swag

swag:
	swag init -d ./cmd/server -o ./docs

.PHONY: dev
dev:
	air