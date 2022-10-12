package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecases usecases.Usecase
}

func InitialiseUserController(usecases usecases.Usecase) UserController {
	return UserController{usecases: usecases}
}

func GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "UserController - GetUsers"})
}

func GetUserById(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "UserController - GetUserById"})
}
