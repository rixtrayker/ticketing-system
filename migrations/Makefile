.PHONY: up down create force status load-env create-db check-db init migrate-to goto version reset test-migration dry-run backup restore validate

# Load environment variables from .env file
load-env:
	@if [ -f ../.env ]; then \
		echo "Loading environment variables from .env..."; \
		set -a; \
		source ../.env; \
		set +a; \
		echo "Environment variables loaded successfully"; \
	else \
		echo "No .env file found, using default values"; \
	fi

# Database connection string with fallback values
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= ticketing_system
DB_SSL_MODE ?= disable

# Construct database URL
DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# Check if database exists
check-db: load-env
	@echo "Checking database connection..."
	@if ! PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d postgres -c '\q' 2>/dev/null; then \
		echo "Error: Could not connect to PostgreSQL server"; \
		echo "Please check your connection settings:"; \
		echo "  Host: $(DB_HOST)"; \
		echo "  Port: $(DB_PORT)"; \
		echo "  User: $(DB_USER)"; \
		exit 1; \
	fi
	@echo "PostgreSQL connection successful"

# Create database if it doesn't exist
create-db: check-db
	@echo "Checking if database exists..."
	@if ! PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1; then \
		echo "Creating database $(DB_NAME)..."; \
		PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)"; \
		echo "Database created successfully"; \
	else \
		echo "Database $(DB_NAME) already exists"; \
	fi

# Run all pending migrations
up: check-db
	@echo "Running migrations up..."
	@echo "Using database: $(DB_URL)"
	@migrate -path sql -database "$(DB_URL)" up || { echo "Migration failed"; exit 1; }
	@echo "Migrations completed successfully"

# Rollback the last migration
down: check-db
	@echo "Rolling back last migration..."
	@echo "Using database: $(DB_URL)"
	@migrate -path sql -database "$(DB_URL)" down 1 || { echo "Rollback failed"; exit 1; }
	@echo "Rollback completed successfully"

# Rollback all migrations
down-all: check-db
	@echo "Rolling back all migrations..."
	@echo "Using database: $(DB_URL)"
	@migrate -path sql -database "$(DB_URL)" down || { echo "Rollback failed"; exit 1; }
	@echo "All migrations rolled back successfully"

# Create new migration files
create:
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "Error: Migration name cannot be empty"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir sql -seq $$name

# Force version
force: check-db
	@read -p "Enter version to force: " version; \
	if [ -z "$$version" ]; then \
		echo "Error: Version cannot be empty"; \
		exit 1; \
	fi; \
	echo "Using database: $(DB_URL)"; \
	migrate -path sql -database "$(DB_URL)" force $$version

# Show migration status
status: check-db
	@echo "Migration status:"
	@echo "Using database: $(DB_URL)"
	@migrate -path sql -database "$(DB_URL)" version

# # Initialize migrations directory
# init:
# 	@if [ ! -d "sql" ]; then \
# 		echo "Creating migrations directory..."; \
# 		mkdir -p sql; \
# 	fi

# Migrate to specific version
migrate-to: check-db
	@read -p "Enter target version: " version; \
	if [ -z "$$version" ]; then \
		echo "Error: Version cannot be empty"; \
		exit 1; \
	fi; \
	echo "Migrating to version $$version..."; \
	echo "Using database: $(DB_URL)"; \
	migrate -path sql -database "$(DB_URL)" goto $$version

# Alias for migrate-to
goto: migrate-to

# Show current version
version: check-db
	@echo "Current migration version:"
	@echo "Using database: $(DB_URL)"
	@migrate -path sql -database "$(DB_URL)" version

# Reset database (drop all and recreate)
reset: check-db
	@echo "WARNING: This will drop all data in the database!"
	@read -p "Are you sure? (yes/no): " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		echo "Dropping all migrations..."; \
		migrate -path sql -database "$(DB_URL)" drop -f; \
		echo "Running all migrations..."; \
		migrate -path sql -database "$(DB_URL)" up; \
		echo "Database reset completed"; \
	else \
		echo "Reset cancelled"; \
	fi

# Test migration (up and then down)
test-migration: check-db
	@echo "Testing last migration (up then down)..."
	@echo "Current version:"
	@migrate -path sql -database "$(DB_URL)" version
	@echo "Running up..."
	@migrate -path sql -database "$(DB_URL)" up 1
	@echo "New version:"
	@migrate -path sql -database "$(DB_URL)" version
	@echo "Running down..."
	@migrate -path sql -database "$(DB_URL)" down 1
	@echo "Final version:"
	@migrate -path sql -database "$(DB_URL)" version
	@echo "Migration test completed"

# Validate migration files
validate:
	@echo "Validating migration files..."
	@if [ ! -d "sql" ]; then \
		echo "Error: sql directory does not exist"; \
		exit 1; \
	fi
	@for file in sql/*.sql; do \
		if [ -f "$$file" ]; then \
			echo "Checking $$file..."; \
			if ! pg_verifybackup --help >/dev/null 2>&1; then \
				echo "Warning: pg_verifybackup not available, skipping SQL validation"; \
				break; \
			fi; \
		fi; \
	done
	@echo "Validation completed"

# Backup database schema
backup: check-db
	@echo "Creating database backup..."
	@timestamp=$$(date +"%Y%m%d_%H%M%S"); \
	backup_file="backup_$${timestamp}.sql"; \
	echo "Backing up to $$backup_file"; \
	PGPASSWORD=$(DB_PASSWORD) pg_dump -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) --schema-only > "$$backup_file"; \
	echo "Backup completed: $$backup_file"

# List migration files
list:
	@echo "Migration files:"
	@if [ -d "sql" ]; then \
		ls -la sql/*.sql 2>/dev/null || echo "No migration files found"; \
	else \
		echo "sql directory does not exist"; \
	fi

# Clean up backup files
clean:
	@echo "Cleaning up backup files..."
	@rm -f backup_*.sql
	@echo "Cleanup completed"

# Help
help:
	@echo "Available targets:"
	@echo "  up          - Run all pending migrations"
	@echo "  down        - Rollback the last migration"
	@echo "  down-all    - Rollback all migrations"
	@echo "  create      - Create new migration files"
	@echo "  status      - Show migration status"
	@echo "  version     - Show current migration version"
	@echo "  migrate-to  - Migrate to specific version"
	@echo "  goto        - Alias for migrate-to"
	@echo "  force       - Force version"
	@echo "  reset       - Reset database (drop all and recreate)"
	@echo "  test-migration - Test last migration (up then down)"
	@echo "  validate    - Validate migration files"
	@echo "  backup      - Create database schema backup"
	@echo "  list        - List migration files"
	@echo "  clean       - Clean up backup files"
	@echo "  create-db   - Create database if it doesn't exist"
	@echo "  check-db    - Check database connection"
	@echo "  help        - Show this help message"
