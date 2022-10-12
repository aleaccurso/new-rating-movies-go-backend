package controllers

import (
	"net/http"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecases usecases.Usecase
}

func InitialiseUserController(usecases usecases.Usecase) UserController {
	return UserController{usecases: usecases}
}

func (controller UserController) GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "UserController - GetUsers"})
}

func (controller UserController) GetUserById(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "UserController - GetUserById"})
}
