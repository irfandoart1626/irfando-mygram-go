package main

import (
	"github.com/gin-gonic/gin"
    "mygram/controllers"
    "mygram/config"
)

func main() {
	// Inisialisasi router
	router := gin.Default()

	// Koneksi ke database PostgreSQL
	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Routes
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)
	commentController := controllers.NewCommentController(db)
	socialMediaController := controllers.NewSocialMediaController(db)

	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)

	// Authenticated routes
	auth := router.Group("/api")
	auth.Use(authMiddleware()) // Middleware untuk autentikasi JWT
	{
		auth.GET("/photos", photoController.GetPhotos)
		auth.POST("/photos", photoController.CreatePhoto)
		auth.GET("/comments", commentController.GetComments)
		auth.POST("/comments", commentController.CreateComment)
		auth.GET("/socialmedia", socialMediaController.GetSocialMedia)
	}

	// Authorized routes
	auth.Use(authorizationMiddleware()) // Middleware untuk autorisasi
	{
		auth.PUT("/photos/:id", photoController.UpdatePhoto)
		auth.DELETE("/photos/:id", photoController.DeletePhoto)
		auth.PUT("/comments/:id", commentController.UpdateComment)
		auth.DELETE("/comments/:id", commentController.DeleteComment)
		auth.PUT("/socialmedia/:id", socialMediaController.UpdateSocialMedia)
		auth.DELETE("/socialmedia/:id", socialMediaController.DeleteSocialMedia)
	}

	// Jalankan server pada port tertentu
	router.Run(":8080")
}
