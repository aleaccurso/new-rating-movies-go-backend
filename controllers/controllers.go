package controllers

import "github.com/gin-gonic/gin"

type Controllers struct {
	UserController UserController
}

type UserController interface {
	GetUsers(*gin.Context)
	GetUserById(*gin.Context)
}
