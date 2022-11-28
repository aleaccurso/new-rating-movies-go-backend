package usecases

import (
	"context"
	"new-rating-movies-go-backend/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthUsecase interface {
	Register(context context.Context, userDTO dtos.UserReqCreateDTO) (*primitive.ObjectID, error)
	Login(context context.Context, loginReqDTO dtos.LoginReqDTO) (*string, error)
}
