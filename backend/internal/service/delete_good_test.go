package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDeleteGood_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteGood/:id", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.DeleteGood(c)
	})

	// Simulate a valid good ID in the URL
	goodID := 1
	req := httptest.NewRequest(http.MethodDelete, "/deleteGood/"+strconv.Itoa(goodID), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteGood_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteGood/:id", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.DeleteGood(c)
	})

	// Simulate a valid good ID in the URL
	goodID := 1
	req := httptest.NewRequest(http.MethodDelete, "/deleteGood/"+strconv.Itoa(goodID), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestDeleteGood_InvalidID(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteGood/:id", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.DeleteGood(c)
	})

	// Simulate an invalid ID in the URL (non-numeric)
	req := httptest.NewRequest(http.MethodDelete, "/deleteGood/invalidID", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
