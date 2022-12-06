package usecases

import (
	"new-rating-movies-go-backend/dtos"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMovieUsecase interface {
	GetMovies(c *gin.Context, page string, size string) (*dtos.MoviePagingResDTO, error)
	GetMovieById(c *gin.Context, movieId string) (*dtos.MovieResDTO, error)
	CreateMovie(c *gin.Context, reqCreateDTO dtos.MovieReqCreateDTO) (*dtos.MovieResDTO, error)
	// UpdateMovieById(c *gin.Context, movieId string, reqUpdateDTO dtos.MovieReqUpdateDTO) (*dtos.MovieResDTO, error)
	DeleteMovieById(c *gin.Context, movieId string) (*primitive.ObjectID, error)
}
