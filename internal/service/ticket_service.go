package service

import (
	"github.com/google/uuid"
	"github.com/rixtrayker/ticketing-system/internal/models"
	"github.com/rixtrayker/ticketing-system/internal/repository"
)

type TicketService interface {
	CreateTicket(input *CreateTicketInput) (*models.Ticket, error)
	UpdateTicket(id uuid.UUID, input *UpdateTicketInput) (*models.Ticket, error)
	DeleteTicket(id uuid.UUID) error
	GetTicket(id uuid.UUID) (*models.Ticket, error)
	GetTickets(filter *models.TicketFilter) ([]*models.Ticket, error)
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	userRepo   repository.UserRepository
	assetRepo  repository.AssetRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, userRepo repository.UserRepository, assetRepo repository.AssetRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
		assetRepo:  assetRepo,
	}
}

func (s *ticketService) CreateTicket(input *CreateTicketInput) (*models.Ticket, error) {
	ticket := &models.Ticket{
		Title:       input.Title,
		Description: input.Description,
		Status:      models.TicketStatusOpen,
		Priority:    input.Priority,
		CreatedByID: input.CreatedByID,
	}

	if input.AssignedToID != nil {
		ticket.AssignedToID = input.AssignedToID
	}

	if input.AssetID != nil {
		ticket.AssetID = input.AssetID
	}

	err := s.ticketRepo.Create(ticket)
	if err != nil {
		return nil, err
	}

	return s.ticketRepo.GetByID(ticket.ID)
}

func (s *ticketService) UpdateTicket(id uuid.UUID, input *UpdateTicketInput) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		ticket.Title = *input.Title
	}
	if input.Description != nil {
		ticket.Description = *input.Description
	}
	if input.Status != nil {
		ticket.Status = *input.Status
	}
	if input.Priority != nil {
		ticket.Priority = *input.Priority
	}
	if input.AssignedToID != nil {
		ticket.AssignedToID = input.AssignedToID
	}

	err = s.ticketRepo.Update(ticket)
	if err != nil {
		return nil, err
	}

	return s.ticketRepo.GetByID(id)
}

func (s *ticketService) DeleteTicket(id uuid.UUID) error {
	return s.ticketRepo.Delete(id)
}

func (s *ticketService) GetTicket(id uuid.UUID) (*models.Ticket, error) {
	return s.ticketRepo.GetByID(id)
}

func (s *ticketService) GetTickets(filter *models.TicketFilter) ([]*models.Ticket, error) {
	return s.ticketRepo.GetAll(filter)
}

// Input types for service layer
type CreateTicketInput struct {
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	Priority     models.TicketPriority `json:"priority"`
	CreatedByID  uuid.UUID             `json:"createdById"`
	AssignedToID *uuid.UUID            `json:"assignedToId,omitempty"`
	AssetID      *uuid.UUID            `json:"assetId,omitempty"`
}

type UpdateTicketInput struct {
	Title        *string                `json:"title,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Status       *models.TicketStatus   `json:"status,omitempty"`
	Priority     *models.TicketPriority `json:"priority,omitempty"`
	AssignedToID *uuid.UUID             `json:"assignedToId,omitempty"`
} 