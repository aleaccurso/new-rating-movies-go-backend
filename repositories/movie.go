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

type MovieRepository struct {
	database *database.Database
}

func InitialiseMovieRepository(db *database.Database) MovieRepository {
	return MovieRepository{
		database: db,
	}
}

func (repository MovieRepository) GetMovies(context context.Context, page int, size int) ([]models.Movie, error) {
	var movies []models.Movie

	limit := int64(size)
	skip := int64(page * size)
	paginator := options.FindOptions{Limit: &limit, Skip: &skip}

	cursor, err := repository.database.Movies.Find(context, bson.M{}, &paginator)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context) {
		//Create a value into which the single document can be decoded
		var movie models.Movie
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cursor.Close(context)

	return movies, nil
}

func (repository MovieRepository) GetMovieById(context context.Context, movieId primitive.ObjectID) (*models.Movie, error) {
	var movie models.Movie

	err := repository.database.Movies.FindOne(context, bson.M{"_id": movieId}).Decode(&movie)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(constants.RESOURCE_NOT_FOUND + "movie")
	}
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &movie, nil
}

func (repository MovieRepository) ModifyMovieById(context context.Context, movie models.Movie) error {

	movie.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"modified_at": movie.UpdatedAt,
		},
	}

	result, err := repository.database.Movies.UpdateOne(context, bson.M{"_id": movie.Id}, update)
	if err == mongo.ErrNoDocuments {
		return errors.New(constants.RESOURCE_NOT_FOUND + "movie")
	}
	if err != nil {
		return errors.New(constants.SERVER_ERROR)
	}

	if result.MatchedCount != 1 && result.ModifiedCount != 1 {
		return errors.New("something went wrong during the update")
	}

	return nil
}

func (repository MovieRepository) DeleteMovieById(context context.Context, movieId primitive.ObjectID) error {

	result, err := repository.database.Movies.DeleteOne(context, bson.M{"_id": movieId})
	if err == mongo.ErrNoDocuments {
		return errors.New(constants.RESOURCE_NOT_FOUND + "movie")
	}
	if err != nil {
		return errors.New(constants.SERVER_ERROR)
	}

	if result.DeletedCount == 0 {
		return errors.New("couldn't to delete the movie")
	}

	return nil
}

func (repository MovieRepository) CountMovies(context context.Context) (*int64, error) {
	count, err := repository.database.Movies.CountDocuments(context, bson.M{})
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(constants.RESOURCE_NOT_FOUND + "user")
	}
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &count, nil
}