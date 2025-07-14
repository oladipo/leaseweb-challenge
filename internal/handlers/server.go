package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oladipo/leaseweb-challenge/internal/repository"
)

type ServerHandler struct {
	repo *repository.ServerRepository
}

// NewServerHandler creates a new ServerHandler
func NewServerHandler(repo *repository.ServerRepository) *ServerHandler {
	return &ServerHandler{repo: repo}
}

func (h *ServerHandler) GetServers(c *gin.Context) {
	// This function will return all servers
	// For now, we will return a placeholder response

	servers, err := h.repo.GetAllServers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve servers"})
		return
	}

	c.JSON(200, servers)
}

func (h *ServerHandler) FilterServers(c *gin.Context) {
	// This function will filter servers based on criteria
	// For now, we will return a placeholder response
	c.JSON(200, gin.H{
		"message": "Filtered list of servers",
	})
}
