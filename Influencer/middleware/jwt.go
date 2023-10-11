package middleware

import (
	"influence-hub-influencer/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(influencer models.Influencer, secretKey string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = influencer.Email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expiration time (adjust as needed)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
