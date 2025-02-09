package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nextnoetics.com/backend/db"
	"nextnoetics.com/backend/models"
)

// ✅ GET All Tasks for a User & Organization
func GetTasks(c *gin.Context) {
	orgID := c.Query("organization_id")
	userID := c.Query("user_id")

	var tasks []models.Task
	query := db.DB.Where("organization_id = ?", orgID)

	// ✅ Filter by user if provided
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	query.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
