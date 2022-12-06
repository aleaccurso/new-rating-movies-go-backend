package usecases

import (
	"context"
	"errors"
	"math"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieUsecase struct {
	repository repositories.Repository
}

func InitialiseMovieUsecase(repository repositories.Repository) MovieUsecase {
	return MovieUsecase{
		repository: repository,
	}
}

func (usecase MovieUsecase) GetMovies(c *gin.Context, page string, size string) (*dtos.MoviePagingResDTO, error) {
	ctx := context.TODO()

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "page")
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "size")
	}

	movies, err := usecase.repository.MovieRepository.GetMovies(ctx, pageInt, sizeInt)
	if err != nil {
		return nil, err
	}

	moviesDTos := mappers.MovieModelsToResDTOs(movies)

	pagingMovies := dtos.MoviePagingResDTO{
		Page:      int8(pageInt),
		Size:      int8(sizeInt),
		NbPages:   int8(math.Ceil(float64(len(moviesDTos)) / float64(sizeInt))),
		NbResults: int16(len(moviesDTos)),
		Data:      moviesDTos,
	}

	return &pagingMovies, nil
}

func (usecase MovieUsecase) GetMovieById(c *gin.Context, movieId string) (*dtos.MovieResDTO, error) {
	return nil, nil
}

func (usecase MovieUsecase) CreateMovie(c *gin.Context, reqCreateDTO dtos.MovieReqCreateDTO) (*dtos.MovieResDTO, error) {
	return nil, nil
}

// func (usecase MovieUsecase) UpdateMovieById(c *gin.Context, movieId string, reqUpdateDTO dtos.MovieReqUpdateDTO) (*dtos.MovieResDTO, error)

func (usecase MovieUsecase) DeleteMovieById(c *gin.Context, movieId string) (*primitive.ObjectID, error) {
	return nil, nil
}
