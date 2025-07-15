package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/leaseweb-challenge/internal/repository"
)

type ServerHandler struct {
	repo repository.ServerRepositoryInterface
}

// FilterRequest defines the expected JSON payload for filtering servers
// This can be extended based on the fields available in the Server model
type FilterRequest struct {
	Storage  string `json:"storage,omitempty"`
	RAM      string `json:"ram,omitempty"`
	HDD      string `json:"hdd,omitempty"`
	Location string `json:"location,omitempty"`
}

// NewServerHandler creates a new ServerHandler
func NewServerHandler(repo repository.ServerRepositoryInterface) *ServerHandler {
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

	var filters FilterRequest

	// Bind JSON payload
	if err := c.ShouldBindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Convert to repository filter format
	filterMap := make(map[string]string)
	if filters.Storage != "" {
		filterMap["hdd LIKE ?"] = "%" + filters.Storage + "%"
	}
	if filters.RAM != "" {
		filterMap["ram LIKE ?"] = "%" + filters.RAM + "%"
	}
	if filters.HDD != "" {
		filterMap["hdd LIKE ?"] = "%" + filters.HDD + "%"
	}
	if filters.Location != "" {
		filterMap["location LIKE ?"] = "%" + filters.Location + "%"
	}

	// Get filtered results
	servers, err := h.repo.FilterServers(filterMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve servers",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(servers),
		"data":  servers,
	})
}
