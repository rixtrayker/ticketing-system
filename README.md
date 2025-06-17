# Ticketing System - Proactive Maintenance & Asset Management Platform

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13%2B-blue)](https://www.postgresql.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

![Ticketing System Banner](https://placehold.co/1200x300/6c5ce7/ffffff?text=Ticketing+System&font=inter)

## 📋 Overview

**Ticketing System** is a GraphQL-based backend API designed to revolutionize maintenance operations for a large chain of cafes. It shifts the paradigm from a reactive, costly repair model to a proactive, data-driven maintenance strategy. This platform provides a centralized system for tracking maintenance tickets, managing all physical assets via QR codes, scheduling preventive care, and generating insightful reports for management.

### 🎯 Key Benefits

- **Reduced Downtime**: Proactive maintenance scheduling prevents equipment failures
- **Cost Optimization**: Better inventory management and preventive maintenance reduce repair costs
- **Improved Efficiency**: Streamlined workflows and automated task routing
- **Data-Driven Decisions**: Comprehensive reporting and analytics
- **Enhanced Asset Life**: Regular maintenance extends equipment lifespan

---

## ✨ Core Features

* **GraphQL API**: Modern, type-safe API with real-time capabilities
* **QR-Powered Asset Management**: Every piece of equipment gets a unique QR code for instant identification and access to its full maintenance history
* **Comprehensive Ticketing System**: Full CRUD operations for maintenance tickets with filtering and search
* **Automated Preventive Maintenance**: Scheduled maintenance tracking with status management
* **Spare Parts Inventory**: Tracks parts usage and inventory levels with automatic reordering alerts
* **User Management**: Role-based access control for different user types (Admin, Manager, Technician, Staff)
* **Asset Lifecycle Management**: Complete asset tracking from purchase to decommission
* **Maintenance History**: Detailed maintenance records with parts usage tracking

---

## 🚀 Technology Stack

This project is built with a modern, scalable, and type-safe technology stack.

### Backend
* **Go (Golang)** v1.18+
  * `gqlgen` for GraphQL API generation
  * `gorm` for database ORM with PostgreSQL
  * `google/uuid` for UUID generation
  * `gin` for HTTP routing (GraphQL endpoint)

### Database
* **PostgreSQL** v13+
  * JSONB support for flexible metadata storage
  * Custom ENUM types for strict data validation
  * Full-text search capabilities
  * Robust indexing for performance
  * UUID primary keys for better scalability

### Development Tools
* **Air** for hot reloading during development
* **golang-migrate** for database migrations
* **Docker** for containerization (optional)

---

## 📂 Project Structure

```
ticketing-system/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/                # Private application code
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── db/
│   │   └── db.go           # Database connection and setup
│   ├── graph/              # GraphQL implementation
│   │   ├── generated/      # Auto-generated GraphQL code (gqlgen)
│   │   ├── model/          # GraphQL models (generated)
│   │   ├── resolver.go     # Main resolver struct
│   │   ├── schema.graphqls # GraphQL schema definition
│   │   └── schema.resolvers.go # Resolver implementations
│   ├── models/
│   │   └── models.go       # Database models and enums
│   ├── repository/         # Data access layer
│   │   ├── asset_repository.go
│   │   ├── ticket_repository.go
│   │   └── user_repository.go
│   └── service/            # Business logic layer
│       └── ticket_service.go
├── pkg/                    # Public packages
│   ├── database/
│   │   └── database.go     # Database utilities
│   └── logger/             # Logging utilities
├── migrations/             # Database migrations
│   ├── sql/
│   │   ├── 000001_init_schema.up.sql
│   │   └── 000001_init_schema.down.sql
│   ├── Makefile           # Migration commands
│   └── README.md          # Migration documentation
├── scripts/
│   └── setup-db.sh        # Database setup script
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── gqlgen.yml            # GraphQL code generation config
├── Makefile              # Project commands
└── README.md
```

---

## 🏁 Getting Started

To get the Ticketing System backend running locally, follow these steps.

### Prerequisites

* Go (version 1.18+)
* PostgreSQL (version 13+)
* `golang-migrate` CLI tool
* Docker (optional, for containerized setup)

### 1. Install golang-migrate

**macOS (Homebrew):**
```bash
brew install golang-migrate
```

**Linux/macOS (Binary):**
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

### 2. Database Setup

1. **Create a PostgreSQL database**:
   ```sql
   CREATE DATABASE ticketing_system;
   ```

2. **Set environment variables** (create a `.env` file in the project root):
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=ticketing_system
   DB_SSL_MODE=disable
   PORT=8080
   ```

3. **Run database migrations**:
   ```bash
   cd migrations/
   make up
   ```

4. **Verify migration status**:
   ```bash
   make status
   ```

### 3. Backend Setup

1. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

2. **Generate GraphQL code**:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```

3. **Start the server**:
   ```bash
   go run cmd/server/main.go
   ```

The GraphQL API will be running at `http://localhost:8080` with GraphQL Playground available at `http://localhost:8080/`.

### 4. Development Commands

Use the Makefile for common development tasks:

```bash
# Generate GraphQL code
make generate

# Run the server
make run

# Run with hot reload (requires air)
make dev

# Run database migrations
make migrate-up

# Rollback migrations
make migrate-down

# Build the application
make build
```

### 5. Docker Setup (Optional)

For a containerized setup:

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode
docker-compose up -d

# Stop all services
docker-compose down
```

---

## 📖 API Documentation

The API is fully documented via the GraphQL schema. Once the backend server is running, you can explore the entire API interactively using the **GraphQL Playground** available at `http://localhost:8080/`.

### Core Types

- **Ticket**: Maintenance requests with status, priority, and assignments
- **Asset**: Physical equipment with QR codes and maintenance history
- **User**: System users with role-based permissions
- **MaintenanceSchedule**: Scheduled maintenance tasks
- **MaintenanceRecord**: Completed maintenance activities
- **Part**: Inventory items used in maintenance
- **Comment**: Ticket discussion threads

### Key Operations

#### Queries
- `tickets(filter: TicketFilter)`: Get tickets with optional filtering
- `assets(filter: AssetFilter)`: Get assets with optional filtering
- `users(filter: UserFilter)`: Get users with optional filtering

#### Mutations
- `createTicket(input: CreateTicketInput!)`: Create new maintenance ticket
- `updateTicket(id: ID!, input: UpdateTicketInput!)`: Update existing ticket
- `createAsset(input: CreateAssetInput!)`: Register new asset
- `createUser(input: CreateUserInput!)`: Add new user

---

## 🔄 Development Workflow

1. **Branch Management**
   - `main` - Production-ready code
   - `develop` - Integration branch
   - `feature/*` - New features
   - `bugfix/*` - Bug fixes

2. **Code Style**
   - Follow Go best practices and style guide
   - Use `gofmt` for code formatting
   - Write meaningful commit messages
   - Add comments for complex business logic

3. **Database Changes**
   - Create new migration files for schema changes
   - Test migrations both up and down
   - Update models and regenerate GraphQL code

4. **GraphQL Changes**
   - Update `schema.graphqls` for API changes
   - Regenerate code with `make generate`
   - Implement resolver methods
   - Test with GraphQL Playground

---

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Issues**
   ```bash
   # Check PostgreSQL is running
   pg_ctl status
   
   # Verify connection string in .env
   # Ensure database exists
   ```

2. **GraphQL Generation Errors**
   ```bash
   # Clear generated files
   rm -rf internal/graph/generated/
   rm -rf internal/graph/model/
   
   # Regenerate
   go run github.com/99designs/gqlgen generate
   ```

3. **Migration Issues**
   ```bash
   # Check migration status
   cd migrations/ && make status
   
   # Force migration version (use carefully)
   migrate -path ./sql -database $DATABASE_URL force VERSION
   ```

4. **Import Path Issues**
   ```bash
   # Update module name if needed
   go mod edit -module github.com/your-username/ticketing-system
   
   # Update all import paths accordingly
   ```

### Getting Help

- Check the GraphQL Playground for API exploration
- Review database logs for connection issues
- Use `go run -race` to check for race conditions
- Enable debug logging by setting `LOG_LEVEL=debug`

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Update documentation
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Development Setup

```bash
# Install development dependencies
go install github.com/cosmtrek/air@latest
go install github.com/99designs/gqlgen@latest

# Set up pre-commit hooks (optional)
pre-commit install
```

---

## 🗄️ Database Schema

The system uses PostgreSQL with the following main entities:

- **users**: System users with roles and authentication
- **assets**: Physical equipment with QR codes and metadata
- **tickets**: Maintenance requests and work orders
- **maintenance_schedules**: Planned maintenance activities
- **maintenance_records**: Completed maintenance history
- **parts**: Inventory items and spare parts
- **part_usages**: Parts consumed during maintenance
- **comments**: Ticket discussions and updates

All tables use UUID primary keys and include created_at, updated_at, and deleted_at timestamps for audit trails.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- Built with [gqlgen](https://gqlgen.com/) for GraphQL code generation
- Database migrations powered by [golang-migrate](https://github.com/golang-migrate/migrate)
- ORM functionality provided by [GORM](https://gorm.io/)
- Thanks to all contributors and the Go community
