package controllers

import (
	"errors"
	"net/http"
	"new-rating-movies-go-backend/constants"
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
	title := c.Param("title")
	if title == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"title").Error())
		return
	}

	language := c.Param("language")
	if title == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"language").Error())
		return
	}

	searchResult, err := controller.services.TheMovieDbService.GetSearchResultsFromAPI(c, title, language)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searchResult)
}

func (controller TheMovieDbController) GetInfoFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
