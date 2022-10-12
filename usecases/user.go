package usecases

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
)

type UserUsecase struct {
	repository repositories.Repository
}

func InitialiseUserUsecase(repository repositories.Repository) UserUsecase {
	return UserUsecase{
		repository: repository,
	}
}

func (usecase UserUsecase) GetUsers() ([]dtos.UserResDTO, error) {
	return nil, nil
}

func (usecase UserUsecase) GetUserById() (*dtos.UserResDTO, error) {
	return nil, nil
}
