package controllers

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "/mygram/models"
)

type PhotoController struct {
    DB *sql.DB
}

func NewPhotoController(db *sql.DB) *PhotoController {
    return &PhotoController{DB: db}
}

// CreatePhoto membuat foto baru dalam database
func (c *PhotoController) CreatePhoto(ctx *gin.Context) {
    var photo models.Photo
    if err := ctx.BindJSON(&photo); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err := c.DB.Exec("INSERT INTO photos (title, caption, photo_url, user_id) VALUES ($1, $2, $3, $4)", photo.Title, photo.Caption, photo.PhotoURL, photo.UserID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to create photo"})
        return
    }

    ctx.JSON(200, gin.H{"message": "Photo created successfully"})
}

// GetPhoto mengambil informasi foto berdasarkan ID
func (c *PhotoController) GetPhoto(ctx *gin.Context) {
    // Ambil ID foto dari path parameter
    photoID := ctx.Param("id")

    var photo models.Photo
    err := c.DB.QueryRow("SELECT id, title, caption, photo_url, user_id FROM photos WHERE id = $1", photoID).Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "Photo not found"})
        return
    }

    ctx.JSON(200, gin.H{"photo": photo})
}

// UpdatePhoto mengupdate informasi foto berdasarkan ID
func (c *PhotoController) UpdatePhoto(ctx *gin.Context) {
    // Ambil ID foto dari path parameter
    photoID := ctx.Param("id")

    var photo models.Photo
    if err := ctx.BindJSON(&photo); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err := c.DB.Exec("UPDATE photos SET title=$1, caption=$2, photo_url=$3 WHERE id=$4", photo.Title, photo.Caption, photo.PhotoURL, photoID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to update photo"})
        return
    }

    ctx.JSON(200, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto menghapus foto berdasarkan ID
func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
    // Ambil ID foto dari path parameter
    photoID := ctx.Param("id")

    _, err := c.DB.Exec("DELETE FROM photos WHERE id=$1", photoID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to delete photo"})
        return
    }

    ctx.JSON(200, gin.H{"message": "Photo deleted successfully"})
}
