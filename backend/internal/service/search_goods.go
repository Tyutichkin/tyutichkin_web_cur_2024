package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) SearchGoods(c *gin.Context) {
	var (
		searchGoodRequest models.SearchGoodRequest
	)
	if err := c.BindJSON(&searchGoodRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	goods, err := s.repo.SearchGoods(c, searchGoodRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": goods})
}
