package repositories

import (
	"context"
	"errors"
	"fmt"
	"new-rating-movies-go-backend/database"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
	"reflect"
)

type AuthRepository struct {
	database *database.Database
}

func InitialiseAuthRepository(db *database.Database) AuthRepository {
	return AuthRepository{
		database: db,
	}
}

func (repository AuthRepository) AddUser(context context.Context, user dtos.UserReqCreateDTO) (*models.User, error) {

	result, err := repository.database.Users.InsertOne(context, user)
	if err != nil {
		return nil, errors.New("reposiotry/unable-to-register")
	}

	newID := result.InsertedID

	fmt.Println("InsertOne() newID:", newID)
	fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	return nil, nil
}
