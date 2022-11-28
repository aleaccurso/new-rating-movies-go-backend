package repositories

import (
	"context"
	"errors"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/database"
	"new-rating-movies-go-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	database *database.Database
}

func InitialiseUserRepository(db *database.Database) UserRepository {
	return UserRepository{
		database: db,
	}
}

func (repository UserRepository) GetUsers(context context.Context, page int, size int) ([]models.User, error) {
	var users []models.User

	cursor, err := repository.database.Users.Find(context, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context) {
		//Create a value into which the single document can be decoded
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cursor.Close(context)

	return users, nil
}

func (repository UserRepository) GetUserById(context context.Context, userId primitive.ObjectID) (*models.User, error) {
	var user models.User

	err := repository.database.Users.FindOne(context, bson.D{{Name: "_id", Value: userId}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &user, nil
}

func (repository UserRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {

	var user models.User

	err := repository.database.Users.FindOne(context, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &user, nil
}
