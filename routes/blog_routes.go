package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"nextnoetics.com/backend/db"
	"nextnoetics.com/backend/models"
)

// ✅ GET All Blog Posts
func GetBlogPosts(c *gin.Context) {
	var posts []models.BlogPost
	db.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"blog_posts": posts})
}

// ✅ CREATE Blog Post
func CreateBlogPost(c *gin.Context) {
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	fmt.Println("Received Post:", post)

	result := db.DB.Create(&post)
	if result.Error != nil {
		fmt.Println("Error saving post:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog post added", "blog_post": post})
}

// ✅ UPDATE Blog Post
func UpdateBlogPost(c *gin.Context) {
	id := c.Param("id")
	var post models.BlogPost

	if err := db.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Blog post updated", "blog_post": post})
}

// ✅ DELETE Blog Post
func DeleteBlogPost(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.BlogPost{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted"})
}
