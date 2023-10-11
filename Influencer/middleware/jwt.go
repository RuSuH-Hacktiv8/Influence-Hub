package middleware

import (
	"influence-hub-influencer/models"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(influencer *models.Influencer, secretKey string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = influencer.ID

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
