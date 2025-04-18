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

func TestAddUser_AdminAllowed(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.POST("/addUser", func(c *gin.Context) {
		c.Set("isAdmin", true)
		s.AddUser(c)
	})

	user := models.User{Login: "admin", Password: "1234"}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/addUser", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddUser_NotAdmin(t *testing.T) {
	r := gin.Default()
	s := NewService(&mockRepo{})
	r.POST("/addUser", func(c *gin.Context) {
		c.Set("isAdmin", false)
		s.AddUser(c)
	})

	req := httptest.NewRequest(http.MethodPost, "/addUser", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
