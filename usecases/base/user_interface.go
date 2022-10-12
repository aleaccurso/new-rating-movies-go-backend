package usecases

import "new-rating-movies-go-backend/dtos"

type IUserUsecaseBase interface {
	GetUsers() ([]dtos.UserResDTO, error)
	GetUserById() (*dtos.UserResDTO, error)
}
