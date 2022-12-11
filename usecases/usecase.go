package usecases

import (
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/services"
)

type Usecase struct {
	UsecaseBase
}

func Initialise(repository repositories.Repository, service services.Service) Usecase {
	return Usecase{
		UsecaseBase: UsecaseBase{
			UserUsecase:  InitialiseUserUsecase(repository),
			MovieUsecase: InitialiseMovieUsecase(repository, service),
		},
	}
}
