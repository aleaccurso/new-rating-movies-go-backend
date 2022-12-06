package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func movieTrailerToResDTO(trailer models.MovieTrailer) dtos.LocalMovieTrailerResDTO {
	return dtos.LocalMovieTrailerResDTO{
		Title: trailer.Title,
		Key:   trailer.Key,
		Site:  trailer.Site,
		Type:  trailer.Type,
	}
}

func movieTrailersToDTOs(trailers []models.MovieTrailer) []dtos.LocalMovieTrailerResDTO {
	trailersDTOs := make([]dtos.LocalMovieTrailerResDTO, len(trailers))

	for i, trailer := range trailers {
		trailersDTOs[i] = movieTrailerToResDTO(trailer)
	}

	return trailersDTOs
}
