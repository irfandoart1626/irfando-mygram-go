package controllers

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/username/mygram/models"
)

type UserController struct {
    DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
    return &UserController{DB: db}
}

// CreateUser membuat user baru dalam database
func (c *UserController) CreateUser(ctx *gin.Context) {
    var user models.User
    if err := ctx.BindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err := c.DB.Exec("INSERT INTO users (username, email, password, age) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Age)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    ctx.JSON(200, gin.H{"message": "User created successfully"})
}

// GetUser mengambil informasi user berdasarkan ID
func (c *UserController) GetUser(ctx *gin.Context) {
    // Ambil ID user dari path parameter
    userID := ctx.Param("id")

    var user models.User
    err := c.DB.QueryRow("SELECT id, username, email, age FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email, &user.Age)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(200, gin.H{"user": user})
}

// UpdateUser mengupdate informasi user berdasarkan ID
func (c *UserController) UpdateUser(ctx *gin.Context) {
    // Ambil ID user dari path parameter
    userID := ctx.Param("id")

    var user models.User
    if err := ctx.BindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err := c.DB.Exec("UPDATE users SET username=$1, email=$2, password=$3, age=$4 WHERE id=$5", user.Username, user.Email, user.Password, user.Age, userID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to update user"})
        return
    }

    ctx.JSON(200, gin.H{"message": "User updated successfully"})
}

// DeleteUser menghapus user berdasarkan ID
func (c *UserController) DeleteUser(ctx *gin.Context) {
    // Ambil ID user dari path parameter
    userID := ctx.Param("id")

    _, err := c.DB.Exec("DELETE FROM users WHERE id=$1", userID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to delete user"})
        return
    }

    ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
