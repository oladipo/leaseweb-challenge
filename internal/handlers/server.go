package handlers

import (
	"net/http"
	"strconv"

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

	// Get pagination parameters from query string
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))

	servers, err := h.repo.GetAllServers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": servers})
}

func (h *ServerHandler) FilterServers(c *gin.Context) {
	// This function will filter servers based on criteria
	// For now, we will return a placeholder response
	c.JSON(http.StatusOK, gin.H{
		"message": "Filtered list of servers",
	})
}
