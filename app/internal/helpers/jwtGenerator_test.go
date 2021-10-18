package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateNewAccessToken(t *testing.T) {
	tokenString, err := GenerateNewAccessToken()

	assert.Nil(t, err)

	assert.NotNil(t, tokenString)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	assert.Nil(t, err)

	assert.NotNil(t, token)

	token, err = jwt.Parse("", jwtKeyFunc)

	assert.NotNil(t, err)

}
