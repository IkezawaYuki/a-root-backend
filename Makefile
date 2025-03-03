.PHONY: swag

swag:
	swag init -g cmd/popple/main.go

.PHONY: dev
dev:
	air