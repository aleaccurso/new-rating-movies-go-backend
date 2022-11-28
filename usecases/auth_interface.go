package usecases

import (
	"context"
	"new-rating-movies-go-backend/dtos"
)

type IAuthUsecase interface {
	Register(context context.Context, userDTO dtos.UserReqCreateDTO) (*dtos.UserResDTO, error)
}