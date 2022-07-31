package main

import (
	"golang-auth/controllers"
	"golang-auth/database"
	"golang-auth/middlewares"

	"github.com/gin-gonic/gin" // Юзаем HTTP-фреймворк
)

func main() {
	// Подключение к базе данных
	database.Connect("root:10RTVx10RTVx@tcp(localhost:3306)/jwt_golang_auth?parseTime=true")
	database.Migrate()

	// Создание роутера
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	// Джин позволяет запихнуть необходимые роуты в группы

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
