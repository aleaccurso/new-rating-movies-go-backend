package controllers

import (
	"context"
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

	ctx := context.TODO()

	if err := c.ShouldBindBodyWith(&userReqCreateDTO, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.UNABLE_TO_BIND_BODY).Error())
		return
	}

	newId, err := controller.services.AuthService.Register(ctx, userReqCreateDTO)
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

	ctx := context.TODO()

	token, err := controller.services.AuthService.Login(ctx, loginReqDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (constroller AuthController) Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(c.Writer, &cookie)

	c.IndentedJSON(http.StatusOK, constants.SUCCESS_ACTION+"logout")
}

func (controller AuthController) GetMe(c *gin.Context) {

	ctx := context.TODO()

	userEmail, ok := c.MustGet("user_email").(string)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, errors.New("email not in context"))
		return
	}

	user, err := controller.usecases.UserUsecase.GetUserByEmail(ctx, userEmail)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}
