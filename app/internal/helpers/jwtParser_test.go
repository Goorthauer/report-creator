package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_FullJwtParser(t *testing.T) {

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_, err := jwtKeyFunc(token)
	assert.Nil(t, err)
	signedString, err := token.SignedString([]byte(""))
	assert.Nil(t, err)
	r, _ := http.NewRequest("POST", "localhost", nil)
	r.Header.Set("Authorization", "Bearer "+signedString)
	_, err = ExtractTokenMetadata(r)
	assert.Nil(t, err)

}
