package repositories

import (
	"context"
	"new-rating-movies-go-backend/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthRepository interface {
	AddUser(context context.Context, user dtos.UserReqCreateDTO) (*primitive.ObjectID, error)
}
