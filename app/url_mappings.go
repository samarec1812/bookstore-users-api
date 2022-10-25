package app

import "github.com/samarec1812/bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:users_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
