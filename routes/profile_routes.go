package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nextnoetics.com/backend/db"
	"nextnoetics.com/backend/models"
)

// ✅ GET User Profile by UserID
func GetProfile(c *gin.Context) {
	userID := c.Param("userId")
	var profile models.Profile

	if err := db.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
