package repository

import (
	"time"

	"github.com/oladipo/leaseweb-challenge/internal/models"
	"gorm.io/gorm"
)

type ServerRepository struct {
	db *gorm.DB
}

func NewServerRepository(db *gorm.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

// GetAllServers retrieves all servers with optional pagination
func (r *ServerRepository) GetAllServers(page, limit int) ([]models.Server, error) {

	var servers []models.Server

	query := r.db.Model(&models.Server{})

	//Add pagination if requested
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	// Execute query
	result := query.Find(&servers)
	if result.Error != nil {
		return nil, result.Error
	}

	return servers, nil
}

func (r *ServerRepository) FilterServers(criteria map[string]string) ([]models.Server, error) {
	// This function will filter servers based on criteria
	// For now, we will return a placeholder response
	return []models.Server{
		{ID: "1", Model: "Model A", RAM: "16GB", HDD: "500GB", Location: "US", Price: "$100", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, nil
}
