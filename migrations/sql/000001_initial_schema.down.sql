-- Drop indexes
DROP INDEX IF EXISTS idx_ticket_updates_ticket_id;
DROP INDEX IF EXISTS idx_tickets_created_at;
DROP INDEX IF EXISTS idx_tickets_assigned_to_id;
DROP INDEX IF EXISTS idx_tickets_asset_id;
DROP INDEX IF EXISTS idx_tickets_priority;
DROP INDEX IF EXISTS idx_tickets_status;
DROP INDEX IF EXISTS idx_assets_qr_code_id;
DROP INDEX IF EXISTS idx_assets_category_id;
DROP INDEX IF EXISTS idx_assets_branch_id;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_email;

-- Drop tables in reverse order (respecting foreign key dependencies)
DROP TABLE IF EXISTS daily_reports;
DROP TABLE IF EXISTS ticket_parts_usage;
DROP TABLE IF EXISTS inventory_parts;
DROP TABLE IF EXISTS preventive_maintenance_schedules;
DROP TABLE IF EXISTS ticket_updates;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS asset_categories;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS branches;

-- Drop custom types
DROP TYPE IF EXISTS ticket_priority;
DROP TYPE IF EXISTS ticket_status;
DROP TYPE IF EXISTS user_role;

