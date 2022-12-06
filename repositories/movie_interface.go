package repositories

import (
	"context"
	"new-rating-movies-go-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMovieRepository interface {
	GetMovies(context context.Context, page int, size int) ([]models.Movie, error)
	GetMovieById(context context.Context, movieId primitive.ObjectID) (*models.Movie, error)
	ModifyMovieById(context context.Context, movie models.Movie) error
	DeleteMovieById(context context.Context, movieId primitive.ObjectID) error
}
