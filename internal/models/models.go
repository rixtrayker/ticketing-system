package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Base model with common fields
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// JSONB type for metadata
type JSONB datatypes.JSON

// Ticket represents a maintenance or repair request
type Ticket struct {
	Base
	Title       string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	Status      TicketStatus `gorm:"type:ticket_status;not null"`
	Priority    TicketPriority `gorm:"type:ticket_priority;not null"`
	AssignedToID *uuid.UUID `gorm:"type:uuid"`
	CreatedByID uuid.UUID   `gorm:"type:uuid;not null"`
	AssetID     *uuid.UUID  `gorm:"type:uuid"`
	ResolvedAt  *time.Time

	// Relations
	AssignedTo *User
	CreatedBy  User
	Asset      *Asset
	Comments   []Comment
}

// Asset represents a physical item that needs maintenance
type Asset struct {
	Base
	Name               string      `gorm:"not null"`
	Type              AssetType   `gorm:"type:asset_type;not null"`
	Status            AssetStatus `gorm:"type:asset_status;not null"`
	Location          string      `gorm:"not null"`
	QRCode            string      `gorm:"not null;unique"`
	PurchaseDate      time.Time   `gorm:"not null"`
	LastMaintenanceDate *time.Time
	NextMaintenanceDate *time.Time
	Metadata          JSONB

	// Relations
	MaintenanceHistory []MaintenanceRecord
	Tickets           []Ticket
}

// User represents a system user
type User struct {
	Base
	Email string     `gorm:"not null;unique"`
	Name  string     `gorm:"not null"`
	Role  UserRole   `gorm:"type:user_role;not null"`

	// Relations
	AssignedTickets []Ticket
	CreatedTickets  []Ticket
}

// MaintenanceSchedule represents a planned maintenance activity
type MaintenanceSchedule struct {
	Base
	AssetID      uuid.UUID            `gorm:"type:uuid;not null"`
	Frequency    MaintenanceFrequency `gorm:"type:maintenance_frequency;not null"`
	LastPerformed *time.Time
	NextDue      time.Time           `gorm:"not null"`
	AssignedToID uuid.UUID           `gorm:"type:uuid;not null"`
	Status       MaintenanceStatus   `gorm:"type:maintenance_status;not null"`
	Notes        string

	// Relations
	Asset      Asset
	AssignedTo User
}

// MaintenanceRecord represents a completed maintenance activity
type MaintenanceRecord struct {
	Base
	AssetID       uuid.UUID          `gorm:"type:uuid;not null"`
	PerformedByID uuid.UUID          `gorm:"type:uuid;not null"`
	PerformedAt   time.Time          `gorm:"not null"`
	Type          MaintenanceType    `gorm:"type:maintenance_type;not null"`
	Notes         string

	// Relations
	Asset       Asset
	PerformedBy User
	PartsUsed   []PartUsage
}

// PartUsage represents parts used in a maintenance record
type PartUsage struct {
	Base
	PartID              uuid.UUID `gorm:"type:uuid;not null"`
	MaintenanceRecordID uuid.UUID `gorm:"type:uuid;not null"`
	Quantity            int       `gorm:"not null"`

	// Relations
	Part               Part
	MaintenanceRecord  MaintenanceRecord
}

// Part represents an inventory item
type Part struct {
	Base
	Name           string    `gorm:"not null"`
	Description    string    `gorm:"not null"`
	Quantity       int       `gorm:"not null"`
	MinimumQuantity int      `gorm:"not null"`
	Location       string    `gorm:"not null"`
	LastRestocked  time.Time `gorm:"not null"`
}

// Comment represents a comment on a ticket
type Comment struct {
	Base
	TicketID uuid.UUID `gorm:"type:uuid;not null"`
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	Content  string    `gorm:"not null"`

	// Relations
	Ticket Ticket
	User   User
}

// Enums
type TicketStatus string
type TicketPriority string
type AssetType string
type AssetStatus string
type UserRole string
type MaintenanceFrequency string
type MaintenanceStatus string
type MaintenanceType string

const (
	// TicketStatus
	TicketStatusOpen        TicketStatus = "OPEN"
	TicketStatusInProgress  TicketStatus = "IN_PROGRESS"
	TicketStatusResolved    TicketStatus = "RESOLVED"
	TicketStatusClosed      TicketStatus = "CLOSED"
	TicketStatusCancelled   TicketStatus = "CANCELLED"

	// TicketPriority
	TicketPriorityLow      TicketPriority = "LOW"
	TicketPriorityMedium   TicketPriority = "MEDIUM"
	TicketPriorityHigh     TicketPriority = "HIGH"
	TicketPriorityCritical TicketPriority = "CRITICAL"

	// AssetType
	AssetTypeEquipment  AssetType = "EQUIPMENT"
	AssetTypeFurniture  AssetType = "FURNITURE"
	AssetTypeElectronics AssetType = "ELECTRONICS"
	AssetTypePlumbing   AssetType = "PLUMBING"
	AssetTypeHVAC       AssetType = "HVAC"
	AssetTypeOther      AssetType = "OTHER"

	// AssetStatus
	AssetStatusOperational      AssetStatus = "OPERATIONAL"
	AssetStatusMaintenanceNeeded AssetStatus = "MAINTENANCE_NEEDED"
	AssetStatusOutOfService     AssetStatus = "OUT_OF_SERVICE"
	AssetStatusDecommissioned   AssetStatus = "DECOMMISSIONED"

	// UserRole
	UserRoleAdmin      UserRole = "ADMIN"
	UserRoleManager    UserRole = "MANAGER"
	UserRoleTechnician UserRole = "TECHNICIAN"
	UserRoleStaff      UserRole = "STAFF"

	// MaintenanceFrequency
	MaintenanceFrequencyDaily      MaintenanceFrequency = "DAILY"
	MaintenanceFrequencyWeekly     MaintenanceFrequency = "WEEKLY"
	MaintenanceFrequencyMonthly    MaintenanceFrequency = "MONTHLY"
	MaintenanceFrequencyQuarterly  MaintenanceFrequency = "QUARTERLY"
	MaintenanceFrequencyBiannual   MaintenanceFrequency = "BIANNUAL"
	MaintenanceFrequencyAnnual     MaintenanceFrequency = "ANNUAL"

	// MaintenanceStatus
	MaintenanceStatusScheduled MaintenanceStatus = "SCHEDULED"
	MaintenanceStatusInProgress MaintenanceStatus = "IN_PROGRESS"
	MaintenanceStatusCompleted  MaintenanceStatus = "COMPLETED"
	MaintenanceStatusCancelled  MaintenanceStatus = "CANCELLED"
	MaintenanceStatusOverdue    MaintenanceStatus = "OVERDUE"

	// MaintenanceType
	MaintenanceTypePreventive     MaintenanceType = "PREVENTIVE"
	MaintenanceTypeCorrective     MaintenanceType = "CORRECTIVE"
	MaintenanceTypePredictive     MaintenanceType = "PREDICTIVE"
	MaintenanceTypeConditionBased MaintenanceType = "CONDITION_BASED"
)

// Filter types for repositories
type TicketFilter struct {
	Status       *TicketStatus
	Priority     *TicketPriority
	AssignedToID *uuid.UUID
	CreatedByID  *uuid.UUID
	AssetID      *uuid.UUID
}

type AssetFilter struct {
	Type     *AssetType
	Status   *AssetStatus
	Location *string
}

type UserFilter struct {
	Role *UserRole
}

type MaintenanceScheduleFilter struct {
	AssetID      *uuid.UUID
	AssignedToID *uuid.UUID
	Status       *MaintenanceStatus
} 