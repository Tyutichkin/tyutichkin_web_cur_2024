package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) EditGoodStock(c *gin.Context) {
	var (
		goodStock models.GoodStock
	)
	if err := c.BindJSON(&goodStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := s.repo.EditGoodStock(c, goodStock)
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
