-- 
CREATE TYPE user_role AS ENUM (
    'branch_staff',
    'technician',
    'manager',
    'admin'
);

-- 
CREATE TYPE ticket_status AS ENUM (
    'new',
    'assigned',
    'in_progress',
    'on_hold_parts',
    'resolved',
    'closed',
    'cancelled'
);

-- 
CREATE TYPE ticket_priority AS ENUM (
    'low',
    'medium',
    'high',
    'critical'
);

-- Stores information about each cafe location
CREATE TABLE branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Stores user accounts for all roles
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    branch_id UUID REFERENCES branches(id) ON DELETE SET NULL, -- A user can be associated with a primary branch
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Categories for organizing assets (e.g., 'Espresso Machines', 'HVAC')
CREATE TABLE asset_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL
);

-- All physical equipment across all branches
CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    qr_code_id VARCHAR(255) UNIQUE NOT NULL, -- The unique identifier tied to the physical QR code
    branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    category_id UUID REFERENCES asset_categories(id) ON DELETE SET NULL,
    model_number VARCHAR(100),
    purchase_date DATE,
    warranty_until DATE,
    metadata JSONB, -- For extra fields like manuals, specs, etc.
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- The main table for all maintenance requests
CREATE TABLE tickets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ticket_status NOT NULL DEFAULT 'new',
    priority ticket_priority NOT NULL DEFAULT 'medium',
    asset_id UUID NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    created_by_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_to_id UUID REFERENCES users(id) ON DELETE SET NULL, -- Technician assigned to the ticket
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    resolved_at TIMESTAMPTZ,
    closed_at TIMESTAMPTZ
);

-- A log of all updates, comments, and status changes for a ticket
CREATE TABLE ticket_updates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id), -- User who made the update
    comment TEXT,
    old_status ticket_status,
    new_status ticket_status,
    photo_url TEXT, -- URL to an uploaded photo (e.g., in S3)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Defines recurring maintenance tasks
CREATE TABLE preventive_maintenance_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_category_id UUID NOT NULL REFERENCES asset_categories(id) ON DELETE CASCADE,
    task_description TEXT NOT NULL,
    frequency_days INT NOT NULL, -- e.g., 90 for every 90 days
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Stores all available spare parts
CREATE TABLE inventory_parts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    part_name VARCHAR(255) NOT NULL,
    part_number VARCHAR(100) UNIQUE,
    quantity_on_hand INT NOT NULL DEFAULT 0,
    reorder_level INT NOT NULL DEFAULT 5
);

-- Joins tickets and parts to track which parts were used for a repair
CREATE TABLE ticket_parts_usage (
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    part_id UUID NOT NULL REFERENCES inventory_parts(id) ON DELETE RESTRICT,
    quantity_used INT NOT NULL,
    PRIMARY KEY (ticket_id, part_id)
);

-- Indexes for users
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

-- Indexes for assets
CREATE INDEX idx_assets_branch_id ON assets(branch_id);
CREATE INDEX idx_assets_category_id ON assets(category_id);
CREATE INDEX idx_assets_qr_code_id ON assets(qr_code_id);

-- Indexes for tickets (very important)
CREATE INDEX idx_tickets_status ON tickets(status);
CREATE INDEX idx_tickets_priority ON tickets(priority);
CREATE INDEX idx_tickets_asset_id ON tickets(asset_id);
CREATE INDEX idx_tickets_assigned_to_id ON tickets(assigned_to_id);
CREATE INDEX idx_tickets_created_at ON tickets(created_at DESC);

-- Index for ticket updates
CREATE INDEX idx_ticket_updates_ticket_id ON ticket_updates(ticket_id);

-- This table stores a summary of key metrics for each day.
CREATE TABLE daily_reports (
    report_date DATE PRIMARY KEY,
    tickets_created INT NOT NULL DEFAULT 0,
    tickets_resolved INT NOT NULL DEFAULT 0,
    tickets_closed INT NOT NULL DEFAULT 0,
    active_critical_tickets INT NOT NULL DEFAULT 0,
    -- Stored in minutes for precision
    avg_resolution_time_minutes INT,
    -- A JSONB object to store top issues or other rich data
    summary_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

