package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteUser_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteUser/:id", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.DeleteUser(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/deleteUser/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteUser/:id", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.DeleteUser(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/deleteUser/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestDeleteUser_InvalidID(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.DELETE("/deleteUser/:id", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.DeleteUser(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/deleteUser/not-a-number", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
