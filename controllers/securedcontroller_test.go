package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-auth/service"
	"golang-auth/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupProtectedRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	protected := router.Group("/api/secured").Use(service.Auth())
	protected.GET("/ping", Ping)
	return router
}

func TestPingEndpoint_Success(t *testing.T) {
	router := setupProtectedRouter()
	token, _ := auth.GenerateJWT("email@example.com", "username")

	req, _ := http.NewRequest(http.MethodGet, "/api/secured/ping", nil)
	req.Header.Set("Авторизация", token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "токен валиден")
}

func TestPingEndpoint_NoToken(t *testing.T) {
	router := setupProtectedRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/secured/ping", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "Запрос не содержит авторизационный токен")
}

func TestPingEndpoint_InvalidToken(t *testing.T) {
	router := setupProtectedRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/secured/ping", nil)
	req.Header.Set("Авторизация", "wrong_token_1234")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "не могу преобразовать в JWT")
}

