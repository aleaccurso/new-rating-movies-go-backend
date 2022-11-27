package controllers

import "new-rating-movies-go-backend/usecases"

type Controller struct {
	UserController UserController
	AuthController AuthController
}

func Initialise(usecases usecases.Usecase) Controller {
	return Controller{
		UserController: InitialiseUserController(usecases),
		AuthController: InitialiseAuthController(usecases),
	}
}
