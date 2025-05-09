package token

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var jwtConfig *jwtStruct


type jwtStruct struct {
    jwtLifeTime        uint
    jwtRefreshLifeTime uint
    privateKey         *rsa.PrivateKey
    publicKey          *rsa.PublicKey
}


// Load digunakan untuk memuat konfigurasi JWT saat aplikasi dijalankan
func Load() {
    // Baca file kunci privat dari sistem file
    privateKeyBytes, err := os.ReadFile("private.pem")
    // Jika terjadi kesalahan saat membaca file, log error dan hentikan aplikasi
    if err != nil {
        log.Fatalf("Failed to load private key: %v", err)
    }


    // Parse kunci privat dari file PEM
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
    // Jika terjadi kesalahan saat parsing, log error dan hentikan aplikasi
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }


    // Baca file kunci publik dari sistem file
    publicKeyBytes, err := os.ReadFile("public.pem")
    // Jika terjadi kesalahan saat membaca file, log error dan hentikan aplikasi
    if err != nil {
        log.Fatalf("Failed to load public key: %v", err)
    }
   
    // Parse kunci publik dari file PEM
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
    // Jika terjadi kesalahan saat parsing, log error dan hentikan aplikasi
    if err != nil {
        log.Fatalf("Failed to parse public key: %v", err)
    }


    // Simpan semua konfigurasi ke variabel global jwtConfig
    jwtConfig = &jwtStruct{
        jwtLifeTime:        3600,  // Set durasi access token
        jwtRefreshLifeTime: 24000, // Set durasi refresh token
        publicKey:          publicKey,                // Simpan public key
        privateKey:         privateKey,               // Simpan private key
    }
}
