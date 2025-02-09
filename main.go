package main

import (
	"nextnoetics.com/backend/db"
	"nextnoetics.com/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.Default())

	// ✅ Serve uploaded images
	r.Static("/uploads", "./uploads")

	// ✅ Register Routes
	r.GET("/api/profiles/:userId", routes.GetProfile)
	r.GET("/api/tasks", routes.GetTasks)

	r.GET("/api/blog_posts", routes.GetBlogPosts)
	r.POST("/api/blog_posts", routes.CreateBlogPost)
	r.PUT("/api/blog_posts/:id", routes.UpdateBlogPost)
	r.DELETE("/api/blog_posts/:id", routes.DeleteBlogPost)

	r.POST("/api/upload", routes.UploadImage)

	r.Run(":5000")
}
