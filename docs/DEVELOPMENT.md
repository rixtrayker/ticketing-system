# Development Guide

This guide provides information for developers working on the Ticketing System project.

## Development Environment Setup

1. **Install Required Tools**
   ```bash
   # Install Go (1.21 or higher)
   brew install go  # macOS
   
   # Install Air for hot reloading
   go install github.com/cosmtrek/air@latest
   
   # Install golang-migrate for database migrations
   brew install golang-migrate  # macOS
   ```

2. **Clone and Setup Project**
   ```bash
   git clone https://github.com/yourusername/ticketing-system.git
   cd ticketing-system
   
   # Install dependencies
   go mod tidy
   
   # Copy environment file
   cp .env.example .env
   ```

3. **Database Setup**
   ```bash
   # Start PostgreSQL using Docker
   docker-compose up -d db
   
   # Run migrations
   go run cmd/server/main.go migrate
   ```

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── auth/           # Authentication logic
│   ├── config/         # Configuration management
│   ├── graph/          # GraphQL implementation
│   ├── middleware/     # HTTP middleware
│   ├── models/         # Database models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
├── pkg/
│   └── logger/         # Logging utilities
└── scripts/            # Build and deployment scripts
```

## Development Workflow

### 1. Running the Application

```bash
# Development mode with hot reload
make dev

# Regular mode
make run

# Build binary
make build
```

### 2. Database Migrations

```bash
# Create new migration
migrate create -ext sql -dir migrations/sql -seq migration_name

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Check migration status
make migrate-status
```

### 3. GraphQL Development

1. **Update Schema**
   - Edit `internal/graph/schema.graphqls`
   - Add new types, queries, or mutations

2. **Generate Code**
   ```bash
   make generate
   ```

3. **Implement Resolvers**
   - Add resolver methods in `internal/graph/schema.resolvers.go`
   - Implement business logic in service layer
   - Add data access in repository layer

### 4. Testing

```bash
# Run all tests
make test

# Run specific test
go test ./internal/service -run TestName

# Run with coverage
make test-coverage
```

## Code Style Guidelines

### 1. Go Code Style

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for code formatting
- Maximum line length: 100 characters
- Use meaningful variable and function names
- Add comments for exported functions and types

### 2. GraphQL Schema Style

- Use PascalCase for type names
- Use camelCase for field names
- Add descriptions for types and fields
- Group related types together
- Use input types for mutations

### 3. Database Style

- Use snake_case for table and column names
- Add indexes for frequently queried fields
- Use foreign key constraints
- Add appropriate comments for complex tables

## Common Tasks

### Adding a New Feature

1. Create a new branch
   ```bash
   git checkout -b feature/new-feature
   ```

2. Update GraphQL schema
   ```graphql
   # Add new types and operations
   type NewFeature {
     id: ID!
     name: String!
   }
   ```

3. Generate code
   ```bash
   make generate
   ```

4. Implement resolvers and business logic
   ```go
   // Add resolver methods
   func (r *mutationResolver) CreateNewFeature(ctx context.Context, input model.CreateNewFeatureInput) (*model.NewFeature, error) {
       // Implementation
   }
   ```

5. Add tests
   ```go
   func TestCreateNewFeature(t *testing.T) {
       // Test implementation
   }
   ```

6. Create migration if needed
   ```bash
   migrate create -ext sql -dir migrations/sql -seq add_new_feature
   ```

### Debugging

1. **Enable Debug Logging**
   ```bash
   export LOG_LEVEL=debug
   ```

2. **Using Delve**
   ```bash
   # Start server with debugger
   dlv debug cmd/server/main.go
   
   # Set breakpoints
   (dlv) break internal/service/ticket.go:42
   ```

3. **GraphQL Playground**
   - Access at `http://localhost:8080/query`
   - Use for testing queries and mutations
   - Check request/response in browser dev tools

## Best Practices

### 1. Error Handling

- Use custom error types
- Add context to errors
- Handle errors at appropriate levels
- Log errors with proper context

### 2. Testing

- Write unit tests for business logic
- Add integration tests for API endpoints
- Use table-driven tests
- Mock external dependencies

### 3. Security

- Validate all inputs
- Use parameterized queries
- Implement proper authentication
- Follow principle of least privilege

### 4. Performance

- Use appropriate indexes
- Implement caching where needed
- Optimize database queries
- Use connection pooling

## Troubleshooting

### Common Issues

1. **Database Connection Issues**
   ```bash
   # Check PostgreSQL logs
   docker-compose logs db
   
   # Verify connection string
   echo $DATABASE_URL
   ```

2. **GraphQL Generation Errors**
   ```bash
   # Clear generated files
   rm -rf internal/graph/generated/
   
   # Regenerate
   make generate
   ```

3. **Migration Issues**
   ```bash
   # Check migration status
   make migrate-status
   
   # Force version if needed
   migrate -path ./migrations/sql -database $DATABASE_URL force VERSION
   ```

### Getting Help

- Check the [API Documentation](API.md)
- Review the [README](README.md)
- Search existing issues
- Ask in the project chat 