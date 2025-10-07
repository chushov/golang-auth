ipackage controllers

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

func setupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	database.Instance = db
}

func TestRegisterUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := gin.Default()
	router.POST("/api/user/register", RegisterUser)

	user := models.User{
		Name:     "TestName",
		Username: "TestUser",
		Email:    "testuser@example.com",
		Password: "Password123",
	}
	payload, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "testuser@example.com")
}

func TestRegisterUser_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := gin.Default()
	router.POST("/api/user/register", RegisterUser)

	badPayload := []byte(`{"email": "foo"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(badPayload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
}

