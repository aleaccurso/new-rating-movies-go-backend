package repositories

import (
	"context"
	"errors"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/database"
	"new-rating-movies-go-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	limit := int64(size)
	skip := int64(page * size)
	paginator := options.FindOptions{Limit: &limit, Skip: &skip}

	cursor, err := repository.database.Users.Find(context, bson.M{}, &paginator)
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

	err := repository.database.Users.FindOne(context, bson.M{"_id": userId}).Decode(&user)
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

func (repository UserRepository) ModifyUserById(context context.Context, user models.User) error {

	user.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"updated_at":  user.UpdatedAt,
			"nickname":    user.Nickname,
			"email":       user.Email,
			"is_admin":    user.IsAdmin,
			"language":    user.Language,
			"profile_pic": user.ProfilePic,
		},
	}

	result, err := repository.database.Users.UpdateOne(context, bson.M{"_id": user.Id}, update)
	if err == mongo.ErrNoDocuments {
		return errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return errors.New(constants.SERVER_ERROR)
	}

	if result.MatchedCount != 1 && result.ModifiedCount != 1 {
		return errors.New("something went wrong during the update")
	}

	return nil
}

func (repository UserRepository) DeleteUserById(context context.Context, userId primitive.ObjectID) error {

	result, err := repository.database.Users.DeleteOne(context, bson.M{"_id": userId})
	if err == mongo.ErrNoDocuments {
		return errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return errors.New(constants.SERVER_ERROR)
	}

	if result.DeletedCount == 0 {
		return errors.New("couldn't to delete the user")
	}

	return nil
}

func (repository UserRepository) CountUsers(context context.Context) (*int64, error) {
	count, err := repository.database.Users.CountDocuments(context, bson.M{})
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &count, nil
}
