package controllers

import (
	"errors"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/services"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	usecases usecases.Usecase
	services services.Service
}

func InitialiseAuthController(usecases usecases.Usecase, services services.Service) AuthController {
	return AuthController{usecases: usecases, services: services}
}

func (controller AuthController) Register(c *gin.Context) {

	var userReqCreateDTO dtos.UserReqCreateDTO

	if err := c.ShouldBindBodyWith(&userReqCreateDTO, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.UNABLE_TO_BIND_BODY).Error())
		return
	}

	newId, err := controller.services.AuthService.Register(c, userReqCreateDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, newId)
}

func (controller AuthController) Login(c *gin.Context) {

	var loginReqDTO dtos.LoginReqDTO
	if err := c.ShouldBindJSON(&loginReqDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := controller.services.AuthService.Login(c, loginReqDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (controller AuthController) Logout(c *gin.Context) {
	controller.services.AuthService.Logout(c)
	c.IndentedJSON(http.StatusOK, constants.SUCCESS_ACTION+"logout")
}

func (controller AuthController) GetMe(c *gin.Context) {

	userEmail, ok := c.MustGet("user_email").(string)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, errors.New("email not in context"))
		return
	}

	user, err := controller.usecases.UserUsecase.GetUserByEmail(c, userEmail)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}
