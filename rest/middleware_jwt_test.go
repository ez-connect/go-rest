package rest

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	config := JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    "4L2g<^fN~*=$ES.5nde6e-j4+KzY7A)~",
		Expire:        3600000,
	}

	assert.NoError(t, InitJWTMiddleware(config))

	data := jwt.MapClaims{
		"_id": "123",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	tokenString, err := GetJWTSignedString(data)
	assert.NoError(t, err)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SigningKey), nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, data["_id"], claims["_id"])
}
