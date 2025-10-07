package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-auth/models"
	"golang-auth/controllers"
	"golang-auth/database"
	"golang-auth/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupE2ERouter() *gin.Engine {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	database.Instance = db

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/user/register", controllers.RegisterUser)
	router.POST("/api/token", controllers.GenerateToken)
	protected := router.Group("/api/secured").Use(service.Auth())
	protected.GET("/ping", controllers.Ping)
	return router
}

func TestFullAuthFlow(t *testing.T) {
	router := setupE2ERouter()

	// 1. Регистрация пользователя
	user := models.User{
		Name:     "Name",
		Username: "TestUser",
		Email:    "user@example.com",
		Password: "SecretPass1",
	}
	userPayload, _ := json.Marshal(user)
	reqReg, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(userPayload))
	reqReg.Header.Set("Content-Type", "application/json")
	recReg := httptest.NewRecorder()
	router.ServeHTTP(recReg, reqReg)
	assert.Equal(t, http.StatusCreated, recReg.Code)

	// 2. Получение токена
	tokenPayload := map[string]string{
		"email":    "user@example.com",
		"password": "SecretPass1",
	}
	tokenReqBody, _ := json.Marshal(tokenPayload)
	reqToken, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(tokenReqBody))
	reqToken.Header.Set("Content-Type", "application/json")
	recToken := httptest.NewRecorder()
	router.ServeHTTP(recToken, reqToken)
	assert.Equal(t, http.StatusOK, recToken.Code)
	// Достаём токен — заменить на json unmarshal если нужно
	var resp map[string]string
	_ = json.Unmarshal(recToken.Body.Bytes(), &resp)
	token := resp["token"]
	assert.NotEmpty(t, token)

	// 3. Запрос защищённого endpoint
	reqPing, _ := http.NewRequest(http.MethodGet, "/api/secured/ping", nil)
	reqPing.Header.Set("Авторизация", token)
	recPing := httptest.NewRecorder()
	router.ServeHTTP(recPing, reqPing)
	assert.Equal(t, http.StatusOK, recPing.Code)
	assert.Contains(t, recPing.Body.String(), "токен валиден")
}

