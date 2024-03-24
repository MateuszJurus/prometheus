package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mateuszjurus/prometheus/config"
)

var jwtKey = []byte(config.JWTSecretKey) // Secret key for signing JWT tokens

// Claims struct to embed jwt.StandardClaims and add custom claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken creates a JWT token for a given user
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 24 hours

	// Create the JWT claims, including the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
