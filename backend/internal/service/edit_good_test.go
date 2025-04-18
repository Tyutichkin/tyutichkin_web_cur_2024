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

func TestEditGood_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.PUT("/editGood", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.EditGood(c)
	})

	// Simulate a valid good data
	good := models.Good{ID: 1, Name: "Good", Price: 100}
	body, _ := json.Marshal(good)
	req := httptest.NewRequest(http.MethodPut, "/editGood", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditGood_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.PUT("/editGood", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.EditGood(c)
	})

	// Simulate a valid good data
	good := models.Good{ID: 1, Name: "Good", Price: 100}
	body, _ := json.Marshal(good)
	req := httptest.NewRequest(http.MethodPut, "/editGood", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestEditGood_InvalidData(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.PUT("/editGood", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.EditGood(c)
	})

	// Send invalid JSON data
	req := httptest.NewRequest(http.MethodPut, "/editGood", bytes.NewReader([]byte("invalid data")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
