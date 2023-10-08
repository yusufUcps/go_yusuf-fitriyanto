package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(signKey string, userID uint) string {
	var token = generateToken(signKey, userID)
	if token == "" {
		return ""
	}
	return token
}

func generateToken(signKey string, id uint) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return validToken
}
