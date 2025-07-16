package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	filterMap := make(map[string]interface{})
	// if filters.Storage != "" {
	// 	fmt.Printf("Storage Filter: %s\n", filters.Storage)

	// 	filterMap["hdd LIKE ?"] = "%" + filters.Storage + "%"
	// }

	if filters.Storage != "" {
		storageGB, err := convertToGB(filters.Storage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid storage format",
				"details": err.Error(),
			})
			return
		}

		// Create SQL condition to compare total storage in GB
		filterMap["(CAST(REGEXP_REPLACE(SPLIT_PART(hdd, 'x', 1), '[^0-9]', '', 'g') AS INTEGER) * "+
			"CASE WHEN hdd ILIKE '%TB%' THEN 1024 WHEN hdd ILIKE '%GB%' THEN 1 ELSE 0 END * "+
			"CAST(REGEXP_REPLACE(SPLIT_PART(SPLIT_PART(hdd, 'x', 2), 'SATA', 1), '[^0-9]', '', 'g') AS INTEGER)) <= ?"] = int(storageGB)
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

// Helper function to convert storage string to GB
func convertToGB(storage string) (float64, error) {
	// Remove any spaces
	storage = strings.TrimSpace(storage)

	// Extract number and unit
	var value float64
	var unit string

	_, err := fmt.Sscanf(storage, "%f%s", &value, &unit)
	if err != nil {
		return 0, err
	}

	// Convert to GB based on unit
	switch strings.ToUpper(unit) {
	case "TB":
		return value * 1024, nil
	case "GB":
		return value, nil
	default:
		return 0, fmt.Errorf("unsupported unit: %s", unit)
	}
}
