package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func movieLocalInfoToResDTO(localInfo models.MovieLocalInfo) dtos.LocalMovieInfoResDTO {
	return dtos.LocalMovieInfoResDTO{
		PosterPath: localInfo.PosterPath,
		Title:      localInfo.Title,
		Overview:   localInfo.Overview,
		Trailers:   movieTrailersToDTOs(localInfo.Trailers),
	}
}
