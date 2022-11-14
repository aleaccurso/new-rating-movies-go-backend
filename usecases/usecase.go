package usecases

import "new-rating-movies-go-backend/repositories"

type Usecase struct {
	UserUsecase IUserUsecase
}

func Initialise(repository repositories.Repository) Usecase {
	return Usecase{
		UserUsecase: InitialiseUserUsecase(repository),
	}
}
