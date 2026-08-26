include .env
export

MIGRATE := migrate
MIGRATIONS := sql/migrations

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS) -database "$(DATABASE_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS) -database "$(DATABASE_URL)" down

migrate-create:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS) -seq $(name)

seed:
	go run cmd/seed/main.go

swagger:
	swag init -g cmd/api/main.go --output docs

up:
	sudo systemctl start postgresql
