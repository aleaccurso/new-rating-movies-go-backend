package controllers

import "new-rating-movies-go-backend/usecases"

type AuthController struct {
	usecases usecases.Usecase
}

func InitialiseAuthController(usecases usecases.Usecase) AuthController {
	return AuthController{usecases: usecases}
}
