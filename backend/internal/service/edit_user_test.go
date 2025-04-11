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

func TestEditUser_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.PUT("/editUser", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.EditUser(c)
	})

	user := models.User{ID: 1, Login: "new_login", Password: "new_pass"}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/editUser", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditUser_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.PUT("/editUser", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.EditUser(c)
	})

	req := httptest.NewRequest(http.MethodPut, "/editUser", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
