package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(r *http.Request) (*TokenMetadata, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			Expires: expires,
		}, nil
	}

	return nil, err
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
