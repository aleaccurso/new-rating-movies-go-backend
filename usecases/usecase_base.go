package usecases

type UsecaseBase struct {
	UserUsecase IUserUsecase
	AuthUsecase IAuthUsecase
}
