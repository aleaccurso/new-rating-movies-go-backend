package services

type ServiceBase struct {
	AuthService       IAuthService
	TheMovieDbService ITheMovieDbService
}
