package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserAuthToken struct {
    ID       uint64 `json:"id"`
    Email    string `json:"email"`
    Username string `json:"username"`
}
// GenerateToken membuat JWT access token dengan data pengguna di dalamnya
func GenerateToken(data *UserAuthToken) (string, error) {
    // Membuat token baru dengan algoritma RS256 (RSA SHA-256)
    token := jwt.New(jwt.SigningMethodRS256)


    // Mengakses dan mengatur claim token (data yang dikandung token)
    claims := token.Claims.(jwt.MapClaims)
    claims["data"] = data                    // Data pengguna dimasukkan ke dalam claim "data"
    claims["iss"] = "nama_app"           // Issuer: penanda siapa yang membuat token
    claims["iat"] = time.Now().Unix()      // Issued At: waktu token dibuat
    claims["exp"] = time.Now().Add(         // Expired At: waktu token akan kedaluwarsa
        time.Duration(jwtConfig.jwtLifeTime) * time.Second,
    ).Unix()


    // Menandatangani token menggunakan private key RSA dan mengembalikannya dalam bentuk string
    return token.SignedString(jwtConfig.privateKey)
}


// GenerateRefreshToken membuat JWT refresh token hanya dengan ID user
func GenerateRefreshToken(id uint64) (string, error) {
    // Membuat token baru dengan algoritma RS256
    token := jwt.New(jwt.SigningMethodRS256)


    // Mengatur claim token
    claims := token.Claims.(jwt.MapClaims)
    claims["data"] = map[string]uint64{
        "id": id, // Refresh token hanya menyimpan ID user
    }
    claims["iss"] = "nama_app"           // Issuer
    claims["iat"] = time.Now().Unix()              // Issued At
    claims["exp"] = time.Now().Add(         // Expired At
        time.Duration(jwtConfig.jwtLifeTime) * time.Second,
    ).Unix()


    // Menandatangani token dan mengembalikan hasilnya
    return token.SignedString(jwtConfig.privateKey)
}
