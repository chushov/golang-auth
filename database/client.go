package database

import (
	"golang-auth/models"
	"log"

	"gorm.io/driver/mysql" // Драйвер для MySQL, взято для простоты работы на домашней среде
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Не могу подключиться к базе данных")
	}
	log.Println("Подключение к базе данных прошло успешно")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Миграции прошли успешно")
}
