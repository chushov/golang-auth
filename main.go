package main

import (
	"fmt"
	"golang-auth/controllers"
	"golang-auth/database"
	"golang-auth/service"
	"os"

	"github.com/gin-gonic/gin" // TODO: Найти аналог GIN-фреймворку
)

func main() {
	// Формируем строку подключения из переменных окружения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	database.Connect(dsn)
	database.Migrate()

	router := initRouter()

	// Читаем порт из переменной окружения с дефолтным значением 8080
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	err := router.Run(":" + appPort)
	if err != nil {
		return
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(service.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
