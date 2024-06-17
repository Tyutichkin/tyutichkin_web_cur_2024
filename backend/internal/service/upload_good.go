package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) UploadGood(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := file.Open()
	defer f.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var good models.Good

	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&good); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.repo.AddGood(c, good)
	if err != nil {
		switch {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusOK)
}
