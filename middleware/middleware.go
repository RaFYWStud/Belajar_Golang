package middleware

import (
	"fmt"
	"golang-tutorial/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

// Fungsi MiddlewareLogin digunakan untuk memvalidasi JWT token pada header Authorization
func MiddlewareLogin(ctx *gin.Context) {
    // Ambil header Authorization, contoh: "Bearer <token>"
    bearerToken := ctx.GetHeader("Authorization")

    fmt.Println("Authorization Header:", bearerToken)
    // Pisahkan berdasarkan spasi. Seharusnya terdiri dari 2 bagian: "Bearer" dan token-nya
    parts := strings.Split(bearerToken, " ")

    // Validasi format token. Harus berupa "Bearer <token>"
    if len(parts) != 2 || parts[0] != "Bearer" {
        // Jika format tidak valid, kirim respons unauthorized dan hentikan proses
        ctx.JSON(401, gin.H{"error": "Invalid token format"})
        ctx.Abort()
        return
    }

    // Ambil bagian token-nya saja
    tokenStr := parts[1]
    fmt.Println(tokenStr)



    // Panggil fungsi untuk validasi access token (JWT)
    account, err := token.ValidateAccessToken(tokenStr)
    if err != nil {
        // Jika token tidak valid (expired, salah tanda tangan, dll), kirim unauthorized
        fmt.Println("error =", err)
        ctx.JSON(401, gin.H{"error": "Invalid token10"})
        ctx.Abort()
        return
    }

    // Jika token valid, simpan informasi user ke context
    ctx.Set("account", account)


    // Lanjut ke handler berikutnya
    ctx.Next()
}
