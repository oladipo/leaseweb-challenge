package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/leaseweb-challenge/internal/models"
	"github.com/stretchr/testify/assert"
)

// --- Mock Repository ---
type mockServerRepo struct {
	mockGetAllServers func(page, limit int) ([]models.Server, error)
	mockFilterServers func(criteria map[string]string) ([]models.Server, error)
}

func (m *mockServerRepo) GetAllServers(page, limit int) ([]models.Server, error) {
	return m.mockGetAllServers(page, limit)
}

func (m *mockServerRepo) FilterServers(criteria map[string]string) ([]models.Server, error) {
	return m.mockFilterServers(criteria)
}

func setupRouter(handler *ServerHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/servers", handler.GetServers)
	r.POST("/servers/filter", handler.FilterServers)
	return r
}

func TestGetServers_Success(t *testing.T) {
	mockRepo := &mockServerRepo{
		mockGetAllServers: func(page, limit int) ([]models.Server, error) {
			return []models.Server{{ID: "1", Model: "TestServer", RAM: "16GB", HDD: "1TB", Location: "TestLoc", Price: "$100"}}, nil
		},
	}
	handler := NewServerHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/servers", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "TestServer")
}

func TestGetServers_Error(t *testing.T) {
	mockRepo := &mockServerRepo{
		mockGetAllServers: func(page, limit int) ([]models.Server, error) {
			return nil, errors.New("db error")
		},
	}
	handler := NewServerHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/servers", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "db error")
}

func TestFilterServers_Success(t *testing.T) {
	mockRepo := &mockServerRepo{
		mockFilterServers: func(criteria map[string]string) ([]models.Server, error) {
			return []models.Server{{ID: "2", Model: "FilteredServer", RAM: "32GB", HDD: "2TB", Location: "FilterLoc", Price: "$200"}}, nil
		},
	}
	handler := NewServerHandler(mockRepo)
	r := setupRouter(handler)

	filterReq := map[string]string{"ram": "32GB"}
	body, _ := json.Marshal(filterReq)
	req, _ := http.NewRequest("POST", "/servers/filter", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "FilteredServer")
}

func TestFilterServers_BadRequest(t *testing.T) {
	mockRepo := &mockServerRepo{
		mockFilterServers: func(criteria map[string]string) ([]models.Server, error) {
			return nil, nil // Should not be called
		},
	}
	handler := NewServerHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("POST", "/servers/filter", bytes.NewBuffer([]byte("bad json")))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid request body")
}

func TestFilterServers_Error(t *testing.T) {
	mockRepo := &mockServerRepo{
		mockFilterServers: func(criteria map[string]string) ([]models.Server, error) {
			return nil, errors.New("filter error")
		},
	}
	handler := NewServerHandler(mockRepo)
	r := setupRouter(handler)

	filterReq := map[string]string{"ram": "32GB"}
	body, _ := json.Marshal(filterReq)
	req, _ := http.NewRequest("POST", "/servers/filter", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "filter error")
}
