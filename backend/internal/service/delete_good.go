package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
	"strconv"
)

func (s *Service) DeleteGood(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect id"})
	}
	err = s.repo.DeleteGoodByID(c, models.Good{ID: id})
	if err != nil {
		switch {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusOK)
}
