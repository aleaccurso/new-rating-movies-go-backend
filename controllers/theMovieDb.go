package controllers

import (
	"net/http"
	"new-rating-movies-go-backend/services"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
)

type TheMovieDbController struct {
	usecases usecases.Usecase
	services services.Service
}

func InitialiseTheMovieDbController(usecases usecases.Usecase, services services.Service) TheMovieDbController {
	return TheMovieDbController{usecases: usecases, services: services}
}

func (controller TheMovieDbController) GetSearchResultsFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (controller TheMovieDbController) GetInfoFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}