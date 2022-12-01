package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type TheMovieDbService struct{}

func InitialiseTheMovieDbService() TheMovieDbService {
	return TheMovieDbService{}
}

func (service TheMovieDbService) GetSearchResultsFromAPI(c *gin.Context, title string, language string) (*dtos.ApiSearchResDTO, error) {

	if !utils.IsAllowedLanguage(language) {
		return nil, errors.New(constants.BAD_PARAMS + "language")
	}

	theMovieDbAPIURL := os.Getenv("API_URL")
	theMovieDbAPIToken := os.Getenv("API_TOKEN")

	url := theMovieDbAPIURL + "/search/movie?api_key=" + theMovieDbAPIToken + "&query=" + title + "&language=" + language

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResults dtos.ApiSearchResDTO

	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		log.Fatal(err)
	}

	return &searchResults, nil
}

func (service TheMovieDbService) GetInfoFromAPI(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}
