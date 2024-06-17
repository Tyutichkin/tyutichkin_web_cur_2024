package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

func (s *Service) AddGood(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if isAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access is denied"})
		return
	}
	var (
		good models.Good
	)
	if err := c.BindJSON(&good); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := s.setCreatedByUserID(c, &good)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func (s *Service) setCreatedByUserID(c *gin.Context, good *models.Good) (err error) {
	login, _ := c.Get("login")
	user, err := s.repo.GetUserByLogin(c, login.(string))
	if err != nil {
		return err
	}
	good.CreatedByUserID = user.ID
	return
}
