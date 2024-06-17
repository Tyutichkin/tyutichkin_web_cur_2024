package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
	"time"
)

var jwtKey = []byte("secret_key")

func (s *Service) Login(c *gin.Context) {
	var credentials models.Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := s.repo.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var currentUser *models.User
	for _, user := range users {
		if user.Login == credentials.Login && user.Password == credentials.Password {
			currentUser = &user
			break
		}
	}

	if currentUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Login:   credentials.Login,
		IsAdmin: currentUser.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "isAdmin": currentUser.IsAdmin})
}
