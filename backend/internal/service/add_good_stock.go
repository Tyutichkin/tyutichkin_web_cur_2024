package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) AddGoodStock(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	var (
		goodStock models.GoodStock
	)
	if err := c.BindJSON(&goodStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := s.repo.AddGoodStock(c, goodStock)
	if err != nil {
		switch {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusOK)
}
