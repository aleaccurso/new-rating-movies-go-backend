package services

import (
	"net/http"
	"new-rating-movies-go-backend/repositories"

	"github.com/gin-gonic/gin"
)

type TheMovieDbService struct{}

func InitialiseTheMovieDbService(repository repositories.Repository) TheMovieDbService {
	return TheMovieDbService{}
}

func (service TheMovieDbService) GetSearchResultsFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (service TheMovieDbService) GetInfoFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
