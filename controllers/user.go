package controllers

import (
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

	users, err := controller.usecases.UserUsecase.GetUsers(c, page, size)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (controller UserController) GetUserById(c *gin.Context) {

	userId := c.Param("userId")

	user, err := controller.usecases.UserUsecase.GetUserById(c, userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) UpdateUserById(c *gin.Context) {

	userId := c.Param("userId")

	var userReqUpdateDTO dtos.UserReqUpdateDTO
	if err := c.ShouldBindJSON(&userReqUpdateDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := controller.usecases.UserUsecase.UpdateUserById(c, userId, userReqUpdateDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) DeleteUserById(c *gin.Context) {

	userId := c.Param("userId")

	user, err := controller.usecases.UserUsecase.DeleteUserById(c, userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// func (controller UserController) GetUserFavorites(c *gin.Context) {
// 	userId := c.Param("userId")

// 	c.IndentedJSON(http.StatusOK, userId)
// }

func (controller UserController) UpSertUserFavorite(c *gin.Context) {
	userId := c.Param("userId")
	movieDbId := c.Param("movieDbId")

	userDTO, err := controller.usecases.UserUsecase.UpSertUserFavorite(c, userId, movieDbId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}
