package usecases

import (
	"context"
	"errors"
	"math"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/services"
	"new-rating-movies-go-backend/usecases/mappers"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieUsecase struct {
	repositories repositories.Repository
	services     services.Service
}

func InitialiseMovieUsecase(repositories repositories.Repository, services services.Service) MovieUsecase {
	return MovieUsecase{
		repositories: repositories,
		services:     services,
	}
}

func (usecase MovieUsecase) GetMovies(c *gin.Context, page string, size string) (*dtos.MoviePagingResDTO, error) {
	ctx := context.TODO()

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "page")
	}
	if pageInt <= 0 {
		pageInt = 1
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "size")
	}
	if sizeInt <= 0 {
		sizeInt = 8
	}

	moviesCount, err := usecase.repositories.MovieRepository.CountMovies(ctx)
	if err != nil {
		return nil, err
	}

	nbPages := math.Ceil(float64(*moviesCount) / float64(sizeInt))

	if pageInt <= 0 {
		pageInt = 1
	}

	if nbPages == 0 {
		nbPages = 1
	}

	if float64(pageInt) > nbPages {
		pageInt = int(nbPages)
	}

	pagingMovies := dtos.MoviePagingResDTO{
		Page:      int8(pageInt),
		Size:      int8(sizeInt),
		NbPages:   int8(nbPages),
		NbResults: int16(*moviesCount),
	}

	movies, err := usecase.repositories.MovieRepository.GetMovies(ctx, pageInt, sizeInt)
	if err != nil {
		return nil, err
	}

	pagingMovies.Data = mappers.MovieModelsToResDTOs(movies)

	return &pagingMovies, nil
}

func (usecase MovieUsecase) GetMovieById(c *gin.Context, movieId string) (*dtos.MovieResDTO, error) {
	ctx := context.TODO()

	movieUUID, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "movieId")
	}

	movie, err := usecase.repositories.MovieRepository.GetMovieById(ctx, movieUUID)
	if err != nil {
		return nil, err
	}

	movieDTO := mappers.MovieModelToResDTO(*movie)

	return &movieDTO, nil
}

func (usecase MovieUsecase) CreateMovie(c *gin.Context, reqDTO dtos.MovieReqCreateDTO) (*dtos.MovieResDTO, error) {
	ctx := context.TODO()

	movieDbIdStr := strconv.Itoa(int(reqDTO.MovieDbId))

	movieCheck, _ := usecase.repositories.MovieRepository.GetMovieByMoviDBId(ctx, int(reqDTO.MovieDbId))
	if movieCheck != nil {
		return nil, errors.New(constants.RESOURCE_EXISTS + "movie-" + movieDbIdStr)
	}

	// getMovieInfo
	movieInfo, err := usecase.services.TheMovieDbService.GetMovieInfoFromAPI(c, movieDbIdStr)
	if err != nil {
		return nil, err
	}

	// Add movie
	addedMovieId, err := usecase.repositories.MovieRepository.AddMovie(ctx, movieInfo)
	if err != nil {
		return nil, err
	}

	// Retrieve movie info from DB
	addedMovie, err := usecase.repositories.MovieRepository.GetMovieById(ctx, *addedMovieId)
	if err != nil {
		return nil, err
	}

	addedMovieDTO := mappers.MovieModelToResDTO(*addedMovie)

	return &addedMovieDTO, nil
}

// func (usecase MovieUsecase) UpdateMovieById(c *gin.Context, movieId string, reqUpdateDTO dtos.MovieReqUpdateDTO) (*dtos.MovieResDTO, error)

func (usecase MovieUsecase) DeleteMovieById(c *gin.Context, movieId string) (*primitive.ObjectID, error) {

	ctx := context.TODO()

	movieUUID, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "movieId")
	}

	err = usecase.repositories.MovieRepository.DeleteMovieById(ctx, movieUUID)
	if err != nil {
		return nil, err
	}

	return &movieUUID, nil
}
