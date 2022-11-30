package usecases

import "new-rating-movies-go-backend/repositories"

type Usecase struct {
	UsecaseBase
}

func Initialise(repository repositories.Repository) Usecase {
	return Usecase{
		UsecaseBase: UsecaseBase{
			UserUsecase: InitialiseUserUsecase(repository),
		},
	}
}
