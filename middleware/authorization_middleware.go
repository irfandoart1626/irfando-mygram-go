package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Di sini Anda dapat mengimplementasikan logika otorisasi Anda
		// Untuk tujuan demonstrasi, kita hanya akan mencetak pesan
		fmt.Println("Authorization middleware")

		// Panggil handler berikutnya
		c.Next()
	}
}
