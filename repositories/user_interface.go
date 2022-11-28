package repositories

import (
	"context"
	"new-rating-movies-go-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	GetUsers(context context.Context, page int, size int) ([]models.User, error)
	GetUserById(context context.Context, userId primitive.ObjectID) (*models.User, error)
}
