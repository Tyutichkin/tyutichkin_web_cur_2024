package service

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"main/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddGoodStock_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.POST("/addGoodStock", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.AddGoodStock(c)
	})

	goodStock := models.GoodStock{GoodID: 1, GoodCount: 100}
	body, _ := json.Marshal(goodStock)
	req := httptest.NewRequest(http.MethodPost, "/addGoodStock", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddGoodStock_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.POST("/addGoodStock", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.AddGoodStock(c)
	})

	req := httptest.NewRequest(http.MethodPost, "/addGoodStock", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestAddGoodStock_InvalidData(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.POST("/addGoodStock", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.AddGoodStock(c)
	})

	// Send invalid JSON data
	req := httptest.NewRequest(http.MethodPost, "/addGoodStock", bytes.NewReader([]byte("invalid data")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
