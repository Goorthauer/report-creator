package helpers

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateNewAccessToken(t *testing.T) {
	tokenString, err := GenerateNewAccessToken()

	assert.Nil(t, err)

	assert.NotNil(t, tokenString)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	assert.Nil(t, err)

	assert.NotNil(t, token)

	_, err = jwt.Parse("", jwtKeyFunc)

	assert.NotNil(t, err)

}
