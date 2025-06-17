# Database Migrations

This directory contains all database migrations for the Ticketing System project using [golang-migrate](https://github.com/golang-migrate/migrate).

## Prerequisites

- PostgreSQL server running
- `golang-migrate` CLI tool installed
- Environment variables configured (see `.env` file in project root)

### Installing golang-migrate

**macOS (Homebrew):**
```bash
brew install golang-migrate
```

**Linux/macOS (Binary):**
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

## Quick Start

1. **Set up your database connection** by configuring environment variables in the project root `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=ticketing_system
   ```

2. **Create the database** (if it doesn't exist):
   ```bash
   make create-db
   ```

3. **Run all pending migrations**:
   ```bash
   make up
   ```

4. **Check migration status**:
   ```bash
   make status
   ```

## Available Commands

### Basic Operations

| Command | Description |
|---------|-------------|
| `make up` | Run all pending migrations |
| `make down` | Rollback the last migration |
| `make down-all` | Rollback all migrations |
| `make status` | Show current migration status |
| `make version` | Show current migration version |

### Migration Management

| Command | Description |
|---------|-------------|
| `make create` | Create new migration files (prompts for name) |
| `make migrate-to` | Migrate to a specific version |
| `make goto` | Alias for migrate-to |
| `make force` | Force set migration version (use with caution) |

### Database Operations

| Command | Description |
|---------|-------------|
| `make create-db` | Create database if it doesn't exist |
| `make check-db` | Test database connection |
| `make reset` | Drop all data and re-run migrations ⚠️ |
| `make backup` | Create timestamped schema backup |

### Development & Testing

| Command | Description |
|---------|-------------|
| `make test-migration` | Test latest migration (up then down) |
| `make validate` | Validate migration file syntax |
| `make list` | List all migration files |
| `make clean` | Remove backup files |

### Utility

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |

## Migration File Structure

Migrations are stored in the `sql/` directory with the following naming convention:

```
sql/
├── 000001_initial_schema.up.sql    # Creates tables, indexes, etc.
├── 000001_initial_schema.down.sql  # Drops tables, indexes, etc.
├── 000002_add_user_roles.up.sql
├── 000002_add_user_roles.down.sql
└── ...
```

### File Naming Rules

- Files must be prefixed with a sequential number (e.g., `000001_`, `000002_`)
- Must include `.up.sql` for forward migrations
- Must include `.down.sql` for rollback migrations
- Use descriptive names: `add_user_table`, `create_indexes`, etc.

## Database Schema Overview

The current schema includes:

### Core Tables
- **branches** - Cafe locations
- **users** - User accounts (staff, technicians, managers, admins)
- **asset_categories** - Equipment categories
- **assets** - Physical equipment with QR codes
- **tickets** - Maintenance requests
- **ticket_updates** - Ticket history and comments

### Support Tables
- **preventive_maintenance_schedules** - Recurring maintenance tasks
- **inventory_parts** - Spare parts inventory
- **ticket_parts_usage** - Parts used in repairs
- **daily_reports** - Analytics and reporting data

### Custom Types
- **user_role** - `branch_staff`, `technician`, `manager`, `admin`
- **ticket_status** - `new`, `assigned`, `in_progress`, `on_hold_parts`, `resolved`, `closed`, `cancelled`
- **ticket_priority** - `low`, `medium`, `high`, `critical`

## Best Practices

### Creating Migrations

1. **Always create both up and down migrations**:
   ```bash
   make create  # Enter descriptive name when prompted
   ```

2. **Test your migrations locally**:
   ```bash
   make test-migration  # Tests up then down
   ```

3. **Use descriptive names**:
   - ✅ `add_user_email_index`
   - ✅ `create_audit_table` 
   - ❌ `fix_stuff`
   - ❌ `update1`

### Migration Content Guidelines

**DO:**
- Use `IF EXISTS` and `IF NOT EXISTS` where appropriate
- Include proper indexes for foreign keys
- Add comments explaining complex logic
- Use transactions for complex migrations
- Make migrations idempotent when possible

**DON'T:**
- Delete or modify existing migration files
- Include data migrations in schema migrations
- Use database-specific features without consideration
- Forget to test rollbacks (down migrations)

### Example Migration

**000003_add_user_preferences.up.sql:**
```sql
-- Add user preferences table
CREATE TABLE IF NOT EXISTS user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    preference_key VARCHAR(100) NOT NULL,
    preference_value TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, preference_key)
);

-- Add index for efficient lookups
CREATE INDEX IF NOT EXISTS idx_user_preferences_user_id 
    ON user_preferences(user_id);
```

**000003_add_user_preferences.down.sql:**
```sql
-- Drop user preferences table and related objects
DROP INDEX IF EXISTS idx_user_preferences_user_id;
DROP TABLE IF EXISTS user_preferences;
```

## Troubleshooting

### Common Issues

**Migration fails with "dirty" state:**
```bash
# Check current version and force clean state
make version
make force  # Enter the current version number
```

**Connection refused:**
```bash
# Check database connection
make check-db

# Verify environment variables
cat ../.env
```

**Migration file format error:**
- Ensure files follow naming convention: `NNNNNN_name.up.sql` / `NNNNNN_name.down.sql`
- Check for invalid characters in filenames
- Verify files are in the `sql/` directory

### Recovery Procedures

**Reset development database:**
```bash
make reset  # ⚠️ Destroys all data
```

**Rollback problematic migration:**
```bash
make down     # Rollback last migration
make status   # Verify current state
```

**Create backup before major changes:**
```bash
make backup   # Creates timestamped backup file
```

## Environment Variables

The following environment variables can be set in your `.env` file or environment:

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database username |
| `DB_PASSWORD` | `postgres` | Database password |
| `DB_NAME` | `ticketing_system` | Database name |
| `DB_SSL_MODE` | `disable` | SSL mode (disable/require/verify-full) |

## Integration with CI/CD

For automated deployments, you can run migrations as part of your deployment pipeline:

```bash
# In your deployment script
cd migrations/
make check-db    # Verify connection
make backup      # Create backup
make up          # Run pending migrations
```

## Getting Help

Run `make help` to see all available commands, or refer to the [golang-migrate documentation](https://github.com/golang-migrate/migrate) for advanced usage.

