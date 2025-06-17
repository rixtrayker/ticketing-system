-- Drop triggers
DROP TRIGGER IF EXISTS update_comments_updated_at ON comments;
DROP TRIGGER IF EXISTS update_part_usages_updated_at ON part_usages;
DROP TRIGGER IF EXISTS update_parts_updated_at ON parts;
DROP TRIGGER IF EXISTS update_maintenance_records_updated_at ON maintenance_records;
DROP TRIGGER IF EXISTS update_maintenance_schedules_updated_at ON maintenance_schedules;
DROP TRIGGER IF EXISTS update_tickets_updated_at ON tickets;
DROP TRIGGER IF EXISTS update_assets_updated_at ON assets;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS part_usages;
DROP TABLE IF EXISTS parts;
DROP TABLE IF EXISTS maintenance_records;
DROP TABLE IF EXISTS maintenance_schedules;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS users;

-- Drop custom types
DROP TYPE IF EXISTS maintenance_type;
DROP TYPE IF EXISTS maintenance_status;
DROP TYPE IF EXISTS maintenance_frequency;
DROP TYPE IF EXISTS user_role;
DROP TYPE IF EXISTS asset_status;
DROP TYPE IF EXISTS asset_type;
DROP TYPE IF EXISTS ticket_priority;
DROP TYPE IF EXISTS ticket_status; 