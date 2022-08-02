package main

import (
	"golang-auth/controllers"
	"golang-auth/database"
	"golang-auth/service"

	"github.com/gin-gonic/gin" // Юзаем HTTP-фреймворк
)

func main() {
	database.Connect("golang-auth:lD1hSN2fs6mVh1Sj1uC7j_9J@tcp(localhost:3306)/jwt_golang_auth?parseTime=true")
	database.Migrate()

	router := initRouter()
	err := router.Run(":8080")
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
