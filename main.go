package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/users", UserController.GetUsers)
	router.GET("/users/:id", GetUserById)

	router.Run("localhost:8080")
}
