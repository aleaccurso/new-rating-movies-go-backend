package controllers

import "new-rating-movies-go-backend/usecases"

type Controller struct {
	UserController UserController
}

func Initialise(usecases usecases.Usecase) Controller {
	return Controller{
		UserController: InitialiseUserController(usecases),
	}
}
