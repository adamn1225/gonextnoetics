package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"nextnoetics.com/backend/db"
	"nextnoetics.com/backend/models"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.Default())

	// ✅ Serve uploaded images
	r.Static("/uploads", "./uploads")

	// ✅ GET All Blog Posts
	r.GET("/api/blog_posts", func(c *gin.Context) {
		var posts []models.BlogPost
		db.DB.Find(&posts)
		c.JSON(http.StatusOK, gin.H{"blog_posts": posts})
	})

	// ✅ CREATE Blog Post
	r.POST("/api/blog_posts", func(c *gin.Context) {
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

		fmt.Println("Saved Post ID:", post.ID)
		c.JSON(http.StatusCreated, gin.H{"message": "Blog post added", "blog_post": post})
	})

	// ✅ Upload Image
	r.POST("/api/upload", handleUploadImage)

	// ✅ UPDATE Blog Post
	r.PUT("/api/blog_posts/:id", func(c *gin.Context) {
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
	})

	// ✅ DELETE Blog Post
	r.DELETE("/api/blog_posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.DB.Delete(&models.BlogPost{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted"})
	})

	r.Run(":5000")
}

func handleUploadImage(c *gin.Context) {
	// ✅ Ensure uploads directory exists
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// ✅ Save file locally
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// ✅ Return the full image URL
	fileURL := fmt.Sprintf("http://localhost:5000/uploads/%s", file.Filename)
	c.JSON(http.StatusOK, gin.H{"url": fileURL})
}
