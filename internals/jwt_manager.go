package internals

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func NewJWTManager(secretKey string, duration time.Duration) *JWTManager {
	return &JWTManager{SecretKey: secretKey, TokenDuration: duration}
}

func (j *JWTManager) Generate(userID int64, username, name, lastname string, dni int, address string, postalCode, phone int) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		//"store_id":    storeID,
		"name":        name,
		"lastname":    lastname,
		"dni":         dni,
		"address":     address,
		"postal_code": postalCode,
		"phone":       phone,
		"exp":         time.Now().Add(j.TokenDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}
