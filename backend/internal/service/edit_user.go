package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) EditUser(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	var (
		user models.User
	)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := s.repo.EditUser(c, user)
	if err != nil {
		switch {
		//case errors.Is(err, models.ErrBookNotFound):
		//	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Book with id=%v not found", book.ID)})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusOK)
}
