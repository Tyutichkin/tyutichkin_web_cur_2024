package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
	"os"
	"strconv"
)

func (s *Service) DownloadGood(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect id"})
		return
	}

	goods, err := s.repo.SearchGoods(c, models.SearchGoodRequest{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.CreateTemp("", "good-*.json")
	defer os.Remove(file.Name())
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(goods[0]); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JSON"})
		return
	}
	c.FileAttachment(file.Name(), "good.json")
}
