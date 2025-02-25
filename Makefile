DB_MIGRATIONS_DIR :=./database/migrations
DB_URL :="mysql://root:@tcp(localhost:3306)/golang_user_api"
run: #run main.go
	@go run cmd/main.go server
create-migration: #create migration
ifdef table_name
	migrate create -ext sql -dir $(DB_MIGRATIONS_DIR) create_$(table_name)_table
else 
	@echo "Please provide a table_name argument", e.g make create-migration table_name=love
endif
	
migrate-up: #run migration up
	migrate -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR) -verbose  up
migrate-down: #run migration doe
	migrate -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR) down

run-seeder: #run seeders
	@go run cmd/main.go seeders

audit: ## Audit the code base
	go fmt ./...
	go mod tidy -v
	go vet ./...
	go mod verify
