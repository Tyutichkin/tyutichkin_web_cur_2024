package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) GetAllGoods(c *gin.Context) {
	goods, err := s.repo.GetAllGoods(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": goods})
}
