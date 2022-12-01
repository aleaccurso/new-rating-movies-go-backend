package services

import (
	"new-rating-movies-go-backend/dtos"

	"github.com/gin-gonic/gin"
)

type ITheMovieDbService interface {
	GetSearchResultsFromAPI(c *gin.Context, title string, language string) (*dtos.ApiSearchResDTO, error)
}
