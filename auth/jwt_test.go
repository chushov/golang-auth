package auth

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT_Success(t *testing.T) {
	tokenStr, err := GenerateJWT("user@example.com", "username")
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
}

func TestValidateToken_Success(t *testing.T) {
	tokenStr, err := GenerateJWT("user@example.com", "username")
	assert.NoError(t, err)
	err = ValidateToken(tokenStr)
	assert.NoError(t, err)
}

func TestValidateToken_Expired(t *testing.T) {
	// Сгенерируем токен, который уже истёк
	claims := &JWTClaim{
		Email: "user@example.com",
		Username: "username",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	assert.NoError(t, err)
	err = ValidateToken(tokenStr)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "токен просрочен")
}

func TestValidateToken_Invalid(t *testing.T) {
	err := ValidateToken("some_invalid_token_string")
	assert.Error(t, err)
}

