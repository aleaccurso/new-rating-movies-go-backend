package controllers

import (
	"new-rating-movies-go-backend/services"
	"new-rating-movies-go-backend/usecases"
)

type Controller struct {
	UserController UserController
	AuthController AuthController
}

func Initialise(usecases usecases.Usecase, services services.Service) Controller {
	return Controller{
		UserController: InitialiseUserController(usecases),
		AuthController: InitialiseAuthController(usecases, services),
	}
}
