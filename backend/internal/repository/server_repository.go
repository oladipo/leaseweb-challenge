package repository

import (
	"github.com/oladipo/leaseweb-challenge/internal/models"
	"gorm.io/gorm"
)

// ServerRepositoryInterface allows mocking in tests
// (manual mock will be used for now)
//
//go:generate mockgen -destination=server_repository_mock.go -package=repository github.com/oladipo/leaseweb-challenge/internal/repository ServerRepositoryInterface
type ServerRepositoryInterface interface {
	GetAllServers(page, limit int) ([]models.Server, error)
	FilterServers(criteria map[string]interface{}) ([]models.Server, error)
}

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

func (r *ServerRepository) FilterServers(criteria map[string]interface{}) ([]models.Server, error) {

	var servers []models.Server

	query := r.db.Model(&models.Server{})

	// Apply filters
	for key, value := range criteria {
		// query = query.Where(key, value)
		// Handle different value types appropriately
		switch v := value.(type) {
		case float64:
			// For numeric comparisons (like storage size)
			query = query.Where(key, v)
		case string:
			// For string LIKE queries
			query = query.Where(key, v)
		default:
			// Handle other types if needed
			query = query.Where(key, value)
		}
	}

	result := query.Find(&servers)
	if result.Error != nil {
		return nil, result.Error
	}

	return servers, nil
}
