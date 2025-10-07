package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-auth/models"
	"golang-auth/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUser(email, username, password string) {
	user := models.User{Name: "Test", Username: username, Email: email}
	user.HashPassword(password)
	database.Instance.Create(&user)
}

func setupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	database.Instance = db
}

func TestGenerateToken_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	setupUser("user@example.com", "user", "Password123")
	router := gin.Default()
	router.POST("/api/token", GenerateToken)

	payload := map[string]string{
		"email":    "user@example.com",
		"password": "Password123",
	}
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "token")
}

func TestGenerateToken_WrongPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	setupUser("user@example.com", "user", "Password123")
	router := gin.Default()
	router.POST("/api/token", GenerateToken)

	payload := map[string]string{
		"email":    "user@example.com",
		"password": "notright",
	}
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "учетные данные не обнаружены")
}

func TestGenerateToken_NoUserFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := gin.Default()
	router.POST("/api/token", GenerateToken)

	payload := map[string]string{
		"email":    "nonexist@example.com",
		"password": "any",
	}
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
}

func TestGenerateToken_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := gin.Default()
	router.POST("/api/token", GenerateToken)

	badPayload := []byte(`{}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/token", bytes.NewBuffer(badPayload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
}

