package middleware

import (
	"fmt"
	"influence-hub-brand/models"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(brand *models.Brand, secretKey string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = brand.ID

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT")), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
