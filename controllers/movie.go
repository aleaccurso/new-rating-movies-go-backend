package controllers

import (
	"errors"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	usecases usecases.Usecase
}

func InitialiseMovieController(usecases usecases.Usecase) MovieController {
	return MovieController{usecases: usecases}
}

func (controller MovieController) GetMovies(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"page").Error())
		return
	}

	size := c.Query("size")
	if size == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"size").Error())
		return
	}

	movies, err := controller.usecases.MovieUsecase.GetMovies(c, page, size)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}
