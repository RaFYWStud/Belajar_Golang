package token

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// ValidateRefreshToken memvalidasi refresh token dan mengembalikan ID user
func ValidateRefreshToken(token string) (uint64, error) {
    // Parsing token menggunakan publicKey RSA
    parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        // Pastikan algoritma tanda tangan adalah RSA
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // Kembalikan public key untuk verifikasi token
        return jwtConfig.publicKey, nil
    })
    if err != nil {
        return 0, err
    }

    // Ambil klaim jika token valid
    if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
        // Ambil data user dari klaim "data"
        data, valid := claims["data"].(map[string]interface{})
        if !valid {
            return 0, errors.New("invalid token1")
        }

        // Ambil ID user dari data
        id, valid := data["id"]
        if !valid {
            return 0, errors.New("invalid token2")
        }

        // ID dibaca sebagai float64, perlu di-cast ke uint64
        return uint64(id.(float64)), nil
    }

    return 0, errors.New("invalid token3")
}

// ValidateAccessToken memvalidasi access token dan mengembalikan data user
func ValidateAccessToken(token string) (*UserAuthToken, error) {
    // Parsing token menggunakan publicKey RSA
    parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        // Pastikan algoritma tanda tangan adalah RSA
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // Gunakan public key untuk verifikasi
        return jwtConfig.publicKey, nil
    })
    if err != nil {
        return nil, err
    }

    // Ambil klaim jika token valid
    if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
        // Ambil nilai dari klaim "data"
        data, valid := claims["data"]
        if !valid {
            return nil, errors.New("invalid token4")
        }

        // Klaim harus dikonversi ke JSON terlebih dahulu sebelum di-unmarshal ke struct
        jsonData, err := json.Marshal(data)
        if err != nil {
            return nil, errors.New("invalid token5")
        }

        // Unmarshal JSON ke struct UserAuthToken
        var user UserAuthToken
        err = json.Unmarshal(jsonData, &user)
        if err != nil {
            return nil, errors.New("invalid token6")
        }

        return &user, nil
    }

    return nil, errors.New("invalid token7")
}