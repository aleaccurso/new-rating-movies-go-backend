package repositories

import (
	"context"
	"errors"
	"new-rating-movies-go-backend/database"
	"new-rating-movies-go-backend/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository struct {
	database *database.Database
}

func InitialiseAuthRepository(db *database.Database) AuthRepository {
	return AuthRepository{
		database: db,
	}
}

func (repository AuthRepository) AddUser(context context.Context, user dtos.UserReqCreateDTO) (*primitive.ObjectID, error) {

	result, err := repository.database.Users.InsertOne(context, user)
	if err != nil {
		return nil, errors.New("reposiotry/unable-to-register")
	}

	newID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("reposiotry/unable-to-convert-id")
	}

	return &newID, nil
}
