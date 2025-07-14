package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oladipo/leaseweb-challenge/internal/models"
)

type ServerRepository struct {
	db *pgxpool.Pool
}

func NewServerRepository(db *pgxpool.Pool) *ServerRepository {
	return &ServerRepository{db: db}
}
func (r *ServerRepository) GetAllServers() ([]models.Server, error) {
	// This function will return all servers
	// For now, we will return a placeholder response
	return []models.Server{
		{ID: "1", Model: "Model A", RAM: "16GB", HDD: "500GB", Location: "US", Price: "$100", CreatedAt: "2023-01-01", UpdatedAt: "2023-01-02"},
		{ID: "2", Model: "Model B", RAM: "32GB", HDD: "1TB", Location: "EU", Price: "$200", CreatedAt: "2023-01-03", UpdatedAt: "2023-01-04"},
	}, nil
}

func (r *ServerRepository) FilterServers(criteria map[string]string) ([]models.Server, error) {
	// This function will filter servers based on criteria
	// For now, we will return a placeholder response
	return []models.Server{
		{ID: "1", Model: "Model A", RAM: "16GB", HDD: "500GB", Location: "US", Price: "$100", CreatedAt: "2023-01-01", UpdatedAt: "2023-01-02"},
	}, nil
}
