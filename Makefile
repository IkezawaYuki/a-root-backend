.PHONY: swag

swag:
	swag init -d ./cmd/server -o ./docs

.PHONY: dev
dev:
	air

.PHONY: migrate-new
migrate-new:
ifdef name
	sql-migrate new $(name)
	chmod -R 777 migrations
else
	$(error 新しくマイグレーションを作成する際は make migrate-new name=[名前])
endif

.PHONY: migrate-up
migrate-up:
	sql-migrate up
	sql-migrate status


.PHONY: migrate-down
migrate-down:
	sql-migrate down
	sql-migrate status