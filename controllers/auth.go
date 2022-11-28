package controllers

import (
	"context"
	"errors"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	usecases usecases.Usecase
}

func InitialiseAuthController(usecases usecases.Usecase) AuthController {
	return AuthController{usecases: usecases}
}

func (controller AuthController) Register(c *gin.Context) {

	var userReqCreateDTO dtos.UserReqCreateDTO

	ctx := context.TODO()

	if err := c.ShouldBindBodyWith(&userReqCreateDTO, binding.JSON); err == nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.UNABLE_TO_BIND_BODY).Error())
		return
	}

	user, err := controller.usecases.AuthUsecase.Register(ctx, userReqCreateDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
