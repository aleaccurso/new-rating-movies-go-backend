package usecases

import (
	"new-rating-movies-go-backend/dtos"

	"github.com/gin-gonic/gin"
)

type IMovieUsecase interface {
	GetMovies(c *gin.Context, page string, size string) ([]dtos.MovieResDTO, error)
}
