package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/enums"
	"new-rating-movies-go-backend/services/mappers"
	"new-rating-movies-go-backend/utils"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type TheMovieDbService struct{}

func InitialiseTheMovieDbService() TheMovieDbService {
	return TheMovieDbService{}
}

func (service TheMovieDbService) GetSearchResultsFromAPI(c *gin.Context, title string, language string) ([]dtos.ApiSearchMovieDTO, error) {

	if !utils.IsAllowedLanguage(language) {
		return nil, errors.New(constants.BAD_PARAMS + "language")
	}

	theMovieDbAPIURL := os.Getenv("API_URL")
	theMovieDbAPIToken := os.Getenv("API_TOKEN")

	url := theMovieDbAPIURL + "/search/movie?api_key=" + theMovieDbAPIToken + "&query=" + title + "&language=" + language + "&include_adult=false"

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

	return searchResults.Results, nil
}

func (service TheMovieDbService) GetMovieInfoFromAPI(c *gin.Context, movieDbId string) (*dtos.ApiGetMovieInfoResDTO, error) {

	_, err := strconv.ParseInt(movieDbId, 10, 32)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + movieDbId)
	}

	generalMovieInfo, err := service.getGeneralMovieInfo(movieDbId)
	if err != nil {
		return nil, err
	}

	localMovieInfo, err := service.getLocalMovieInfo(movieDbId)
	if err != nil {
		return nil, err
	}

	movieInfo, err := mappers.ApiMovieInfoToMovieResDTO(*generalMovieInfo, *localMovieInfo)
	if err != nil {
		return nil, err
	}

	return movieInfo, nil

}

func (service TheMovieDbService) getGeneralMovieInfo(movieDbId string) (*dtos.ApiGeneralMovieInfoResDTO, error) {

	theMovieDbAPIURL := os.Getenv("API_URL")
	theMovieDbAPIToken := os.Getenv("API_TOKEN")

	url := theMovieDbAPIURL + "/movie/" + string(movieDbId) + "?api_key=" + theMovieDbAPIToken + "&append_to_response=credits&language=en"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var movieInfoResult dtos.ApiGeneralMovieInfoResDTO

	err = json.Unmarshal(body, &movieInfoResult)
	if err != nil {
		log.Fatal(err)
	}

	if len(movieInfoResult.Genres) > 0 {
		movieInfoResult.Genres = service.genresToLowerCase(movieInfoResult.Genres)
	}

	return &movieInfoResult, nil
}

func (service TheMovieDbService) getLocalMovieInfo(movieDbId string) (*map[string]dtos.LocalMovieInfoResDTO, error) {

	allLocalInfo := map[string]dtos.LocalMovieInfoResDTO{}

	allowedLanguages := enums.AllowedLanguages

	for _, language := range allowedLanguages {

		localInfo, err := service.retrieveLocalInfo(movieDbId, language)
		if err != nil {
			return nil, err
		}

		allLocalInfo[language] = *localInfo
	}

	return &allLocalInfo, nil
}

func (service TheMovieDbService) retrieveLocalInfo(movieDbId string, language string) (*dtos.LocalMovieInfoResDTO, error) {

	theMovieDbAPIURL := os.Getenv("API_URL")
	theMovieDbAPIToken := os.Getenv("API_TOKEN")

	url := theMovieDbAPIURL + "/movie/" + string(movieDbId) + "?api_key=" + theMovieDbAPIToken + "&append_to_response=credits&language=" + language

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var movieInfoResult dtos.LocalMovieInfoResDTO

	err = json.Unmarshal(body, &movieInfoResult)
	if err != nil {
		log.Fatal(err)
	}

	movieInfoResult.Trailers, err = service.retrieveLocalMovieTrailers(movieDbId, language)
	if err != nil {
		return nil, err
	}

	return &movieInfoResult, nil
}

func (service TheMovieDbService) retrieveLocalMovieTrailers(movieDbId string, language string) ([]dtos.LocalMovieTrailerResDTO, error) {

	trailersToReturn := []dtos.LocalMovieTrailerResDTO{}

	theMovieDbAPIURL := os.Getenv("API_URL")
	theMovieDbAPIToken := os.Getenv("API_TOKEN")

	url := theMovieDbAPIURL + "/movie/" + string(movieDbId) + "/videos?api_key=" + theMovieDbAPIToken + "&language=" + language

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var movieTrailersResult dtos.ApiMovieVideoResDTO

	err = json.Unmarshal(body, &movieTrailersResult)
	if err != nil {
		log.Fatal(err)
	}

	for _, trailer := range movieTrailersResult.Results {
		if trailer.Site == "YouTube" && slices.Contains(enums.TrailerTypes, trailer.Type) {
			trailersToReturn = append(trailersToReturn, trailer)
		}
	}

	return trailersToReturn, nil
}

func (service TheMovieDbService) genresToLowerCase(genres []dtos.ApiGenreDTO) []dtos.ApiGenreDTO {
	lowerCaseGenres := make([]dtos.ApiGenreDTO, len(genres))
	for i, genre := range genres {
		lowerCaseGenres[i].Name = strings.ToLower(genre.Name)
	}

	return lowerCaseGenres
}
