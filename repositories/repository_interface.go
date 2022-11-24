package repositories

import (
	"context"
	"new-rating-movies-go-backend/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	GetUsers(context context.Context, page int, size int) ([]dtos.UserResDTO, error)
	GetUserById(context context.Context, userId primitive.ObjectID) (*dtos.UserResDTO, error)
}
