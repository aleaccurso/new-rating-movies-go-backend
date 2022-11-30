package usecases

import (
	"context"
	"new-rating-movies-go-backend/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUsecase interface {
	GetUsers(context context.Context, page string, size string) ([]dtos.UserResDTO, error)
	GetUserById(context context.Context, userId string) (*dtos.UserResDTO, error)
	GetUserByEmail(context context.Context, email string) (*dtos.UserResDTO, error)
	UpdateUserById(context context.Context, userId string, reqUpdateDTO dtos.UserReqUpdateDTO) (*dtos.UserResDTO, error)
	DeleteUserById(context context.Context, userId string) (*primitive.ObjectID, error)
}
