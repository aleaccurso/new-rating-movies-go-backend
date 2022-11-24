package usecases

import (
	"context"
	"new-rating-movies-go-backend/dtos"
)

type IUserUsecase interface {
	GetUsers(context context.Context, page string, size string) ([]dtos.UserResDTO, error)
	GetUserById(context context.Context, userId string) (*dtos.UserResDTO, error)
}
