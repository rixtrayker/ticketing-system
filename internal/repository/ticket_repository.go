package repository

import (
	"github.com/google/uuid"
	"github.com/rixtrayker/ticketing-system/internal/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *models.Ticket) error
	GetByID(id uuid.UUID) (*models.Ticket, error)
	GetAll(filter *models.TicketFilter) ([]*models.Ticket, error)
	Update(ticket *models.Ticket) error
	Delete(id uuid.UUID) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) GetByID(id uuid.UUID) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("AssignedTo").Preload("CreatedBy").Preload("Asset").Preload("Comments").First(&ticket, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) GetAll(filter *models.TicketFilter) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	query := r.db.Preload("AssignedTo").Preload("CreatedBy").Preload("Asset")

	if filter != nil {
		if filter.Status != nil {
			query = query.Where("status = ?", *filter.Status)
		}
		if filter.Priority != nil {
			query = query.Where("priority = ?", *filter.Priority)
		}
		if filter.AssignedToID != nil {
			query = query.Where("assigned_to_id = ?", *filter.AssignedToID)
		}
		if filter.CreatedByID != nil {
			query = query.Where("created_by_id = ?", *filter.CreatedByID)
		}
		if filter.AssetID != nil {
			query = query.Where("asset_id = ?", *filter.AssetID)
		}
	}

	err := query.Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) Update(ticket *models.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Ticket{}, id).Error
} 