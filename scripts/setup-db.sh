#!/bin/bash

# Exit on error
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to check if a command exists
command_exists() {
    command -v "$1" &> /dev/null
}

# Function to check PostgreSQL connection
check_postgres_connection() {
    if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c '\q' 2>/dev/null; then
        echo -e "${RED}Could not connect to PostgreSQL. Please check your connection settings.${NC}"
        echo -e "Host: $DB_HOST"
        echo -e "Port: $DB_PORT"
        echo -e "User: $DB_USER"
        exit 1
    fi
}

# Function to load environment variables
load_env() {
    if [ -f ../.env ]; then
        echo -e "${YELLOW}Loading environment variables from .env...${NC}"
        export $(cat ../.env | grep -v '^#' | xargs)
        echo -e "${GREEN}Environment variables loaded successfully${NC}"
    else
        echo -e "${YELLOW}No .env file found, using default values${NC}"
        # Set default values
        export DB_HOST=${DB_HOST:-localhost}
        export DB_PORT=${DB_PORT:-5432}
        export DB_USER=${DB_USER:-postgres}
        export DB_PASSWORD=${DB_PASSWORD:-postgres}
        export DB_NAME=${DB_NAME:-ticketing_system}
        export DB_SSL_MODE=${DB_SSL_MODE:-disable}
    fi
}

echo -e "${YELLOW}Starting database setup...${NC}"

# Check if PostgreSQL is installed
if ! command_exists psql; then
    echo -e "${RED}PostgreSQL is not installed. Please install it first.${NC}"
    echo -e "${YELLOW}Installation instructions:${NC}"
    echo -e "  - macOS: brew install postgresql"
    echo -e "  - Ubuntu: sudo apt-get install postgresql"
    echo -e "  - Windows: Download from https://www.postgresql.org/download/windows/"
    exit 1
fi

# Check if migrate is installed
if ! command_exists migrate; then
    echo -e "${RED}migrate is not installed. Installing...${NC}"
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    if ! command_exists migrate; then
        echo -e "${RED}Failed to install migrate. Please install it manually:${NC}"
        echo -e "${YELLOW}go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest${NC}"
        exit 1
    fi
fi

# Load environment variables
load_env

# Check PostgreSQL connection
echo -e "${YELLOW}Checking PostgreSQL connection...${NC}"
check_postgres_connection
echo -e "${GREEN}PostgreSQL connection successful${NC}"

# Create database
echo -e "${YELLOW}Creating database...${NC}"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1; then
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c "CREATE DATABASE $DB_NAME"
    echo -e "${GREEN}Database created successfully${NC}"
else
    echo -e "${YELLOW}Database already exists${NC}"
fi

# Run migrations
echo -e "${YELLOW}Running migrations...${NC}"
cd migrations
make up

echo -e "${GREEN}Database setup completed successfully!${NC}"
echo -e "${YELLOW}You can now start the application with:${NC}"
echo -e "  make run"