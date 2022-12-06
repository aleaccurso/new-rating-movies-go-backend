package repositories

type RepositoryBase struct {
	UserRepository  IUserRepository
	AuthRepository  IAuthRepository
	MovieRepository IMovieRepository
}
