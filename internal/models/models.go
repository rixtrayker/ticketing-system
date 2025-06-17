package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string
type TicketStatus string
type TicketPriority string

const (
	RoleAdmin      UserRole = "admin"
	RoleManager    UserRole = "manager"
	RoleTechnician UserRole = "technician"
	RoleUser       UserRole = "user"

	StatusNew       TicketStatus = "new"
	StatusInProgress TicketStatus = "in_progress"
	StatusResolved  TicketStatus = "resolved"
	StatusClosed    TicketStatus = "closed"

	PriorityLow    TicketPriority = "low"
	PriorityMedium TicketPriority = "medium"
	PriorityHigh   TicketPriority = "high"
)

type Branch struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	Address   string
	CreatedAt time.Time `gorm:"not null;default:now()"`
	Users     []User
	Assets    []Asset
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FullName     string    `gorm:"not null"`
	Email        string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null"`
	Role         UserRole  `gorm:"not null"`
	BranchID     *uuid.UUID
	Branch       *Branch
	IsActive     bool      `gorm:"not null;default:true"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
}

type AssetCategory struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"not null;unique"`
	Assets    []Asset
}

type Asset struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name          string    `gorm:"not null"`
	QRCodeID      string    `gorm:"not null;unique"`
	BranchID      uuid.UUID `gorm:"not null"`
	Branch        Branch
	CategoryID    *uuid.UUID
	Category      *AssetCategory
	ModelNumber   string
	PurchaseDate  *time.Time
	WarrantyUntil *time.Time
	Metadata      JSONB
	CreatedAt     time.Time `gorm:"not null;default:now()"`
	Tickets       []Ticket
}

type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `gorm:"not null"`
	Description string
	Status      TicketStatus  `gorm:"not null;default:'new'"`
	Priority    TicketPriority `gorm:"not null;default:'medium'"`
	AssetID     uuid.UUID `gorm:"not null"`
	Asset       Asset
	CreatedByID uuid.UUID `gorm:"not null"`
	CreatedBy   User
	AssignedToID *uuid.UUID
	AssignedTo   *User
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	ResolvedAt  *time.Time
	ClosedAt    *time.Time
	Updates     []TicketUpdate
	PartsUsed   []TicketPartsUsage
}

type TicketUpdate struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	TicketID   uuid.UUID `gorm:"not null"`
	Ticket     Ticket
	UserID     uuid.UUID `gorm:"not null"`
	User       User
	Comment    string
	OldStatus  *TicketStatus
	NewStatus  *TicketStatus
	PhotoURL   string
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

type PreventiveMaintenanceSchedule struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AssetCategoryID uuid.UUID `gorm:"not null"`
	Category        AssetCategory
	TaskDescription string    `gorm:"not null"`
	FrequencyDays   int       `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null;default:now()"`
}

type InventoryPart struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PartName      string    `gorm:"not null"`
	PartNumber    string    `gorm:"unique"`
	QuantityOnHand int       `gorm:"not null;default:0"`
	ReorderLevel  int       `gorm:"not null;default:5"`
	TicketsUsed   []TicketPartsUsage
}

type TicketPartsUsage struct {
	TicketID    uuid.UUID `gorm:"primaryKey"`
	Ticket      Ticket
	PartID      uuid.UUID `gorm:"primaryKey"`
	Part        InventoryPart
	QuantityUsed int       `gorm:"not null"`
}

type DailyReport struct {
	ReportDate            time.Time `gorm:"primaryKey"`
	TicketsCreated       int       `gorm:"not null;default:0"`
	TicketsResolved      int       `gorm:"not null;default:0"`
	TicketsClosed        int       `gorm:"not null;default:0"`
	ActiveCriticalTickets int       `gorm:"not null;default:0"`
	AvgResolutionTime    int       // in minutes
	SummaryData          JSONB
	CreatedAt            time.Time `gorm:"not null;default:now()"`
}

type JSONB map[string]interface{} 