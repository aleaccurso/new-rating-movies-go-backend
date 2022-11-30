package controllers

import (
	"context"
	"errors"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecases usecases.Usecase
}

func InitialiseUserController(usecases usecases.Usecase) UserController {
	return UserController{usecases: usecases}
}

func (controller UserController) GetUsers(c *gin.Context) {

	ctx := context.TODO()

	page := c.Query("page")
	if page == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"page").Error())
		return
	}

	size := c.Query("size")
	if size == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"size").Error())
		return
	}

	users, err := controller.usecases.UserUsecase.GetUsers(ctx, page, size)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (controller UserController) GetUserById(c *gin.Context) {

	ctx := context.TODO()

	userId := c.Param("userId")

	user, err := controller.usecases.UserUsecase.GetUserById(ctx, userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) UpdateUserById(c *gin.Context) {

	ctx := context.TODO()

	userId := c.Param("userId")

	var userReqUpdateDTO dtos.UserReqUpdateDTO
	if err := c.ShouldBindJSON(&userReqUpdateDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := controller.usecases.UserUsecase.UpdateUserById(ctx, userId, userReqUpdateDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
