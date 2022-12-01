package services

import "new-rating-movies-go-backend/repositories"

type Service struct {
	ServiceBase
}

func Initialise(repository repositories.Repository) Service {
	return Service{
		ServiceBase: ServiceBase{
			AuthService:       InitialiseAuthService(repository),
			TheMovieDbService: InitialiseTheMovieDbService(),
		},
	}
}
