package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go" // библиотека для работы с JWT
)

var jwtKey = []byte("supersecretkey") // ключ для подписи JWT

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Используем библиотеку жвт-го, читаем доки.
// По сути все на стандартных методах для проверки подписи и проверки валидности времени.

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("не могу преобразовать в JWTClaim")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("токен просрочен")
		return
	}

	return

}
