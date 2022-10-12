package usecases

type IUserUsecase interface {
	GetUsers(context *contexts.Context, campaignId string, editUserDTO dtos.EditUserDTO) (*dtos.UserDTO, error)
	GetUserById(context *contexts.Context) (*dtos.UserDTO, error)
}
