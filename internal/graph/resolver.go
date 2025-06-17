package graph

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

// Helper function to convert string ID to UUID
func stringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

// Helper function to convert UUID to string
func uuidToString(id uuid.UUID) string {
	return id.String()
} 