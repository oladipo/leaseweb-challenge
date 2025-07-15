package repository

import (
	"fmt"

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

	var servers []models.Server

	query := r.db.Model(&models.Server{})

	fmt.Printf("Applying filters: %v", criteria)
	// Apply filters
	for key, value := range criteria {
		query = query.Where(key, value)
	}

	result := query.Find(&servers)
	if result.Error != nil {
		return nil, result.Error
	}

	return servers, nil
}
