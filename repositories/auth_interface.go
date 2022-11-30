package repositories

import (
	"context"
	"new-rating-movies-go-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthRepository interface {
	AddUser(context context.Context, user models.User) (*primitive.ObjectID, error)
}
