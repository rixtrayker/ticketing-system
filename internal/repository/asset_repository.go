package repository

import (
	"github.com/google/uuid"
	"github.com/rixtrayker/ticketing-system/internal/models"
	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(asset *models.Asset) error
	GetByID(id uuid.UUID) (*models.Asset, error)
	GetAll(filter *models.AssetFilter) ([]*models.Asset, error)
	Update(asset *models.Asset) error
	Delete(id uuid.UUID) error
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepository) GetByID(id uuid.UUID) (*models.Asset, error) {
	var asset models.Asset
	err := r.db.Preload("MaintenanceHistory").Preload("Tickets").First(&asset, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *assetRepository) GetAll(filter *models.AssetFilter) ([]*models.Asset, error) {
	var assets []*models.Asset
	query := r.db.Model(&models.Asset{})

	if filter != nil {
		if filter.Type != nil {
			query = query.Where("type = ?", *filter.Type)
		}
		if filter.Status != nil {
			query = query.Where("status = ?", *filter.Status)
		}
		if filter.Location != nil {
			query = query.Where("location ILIKE ?", "%"+*filter.Location+"%")
		}
	}

	err := query.Find(&assets).Error
	return assets, err
}

func (r *assetRepository) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Asset{}, id).Error
} 