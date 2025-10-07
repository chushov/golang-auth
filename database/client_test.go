package database

import (
	"testing"
	"golang-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

func TestConnect_Success(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestMigrate_UserTable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	Instance = db

	Migrate()
	hasTable := db.Migrator().HasTable(&models.User{})
	assert.True(t, hasTable)
}

