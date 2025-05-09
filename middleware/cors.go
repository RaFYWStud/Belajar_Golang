package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Membuat semua domain bisa mengakses API,
		// * (wildcard) Yang memberi akses ke semua domain, ini bisa diganti dengan url frontend

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		// Mengizinkan metode HTTP yang digunakan oleh frontend

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Mengizinkan header yang digunakan oleh frontend

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// Mengizinkan credential seperti cookie, authorization header, atau TLS client certificates

		// Jika permintaan adalah preflight (OPTIONS), langsung kirim respons 204 (No Content)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Jika bukan preflight, lanjutkan ke handler berikutnya
		c.Next()
	}
}
