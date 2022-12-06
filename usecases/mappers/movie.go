package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func MovieModelToResDTO(movie models.Movie) dtos.MovieResDTO {
	return dtos.MovieResDTO{
		Id:          movie.Id,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		MovieDbId:   movie.MovieDbId,
		ReleaseDate: movie.ReleaseDate,
		Director:    movie.Director,
		Casting:     movie.Casting,
		VoteAverage: float32(movie.VoteAverage),
		VoteCount:   float32(movie.VoteCount),
		Genre:       movie.Genre,
		En:          movieLocalInfoToResDTO(movie.En),
		Fr:          movieLocalInfoToResDTO(movie.Fr),
		It:          movieLocalInfoToResDTO(movie.It),
		Nl:          movieLocalInfoToResDTO(movie.Nl),
	}
}

func MovieModelsToResDTOs(movies []models.Movie) []dtos.MovieResDTO {
	moviesDTOs := make([]dtos.MovieResDTO, len(movies))

	for i, movie := range movies {
		moviesDTOs[i] = MovieModelToResDTO(movie)
	}

	return moviesDTOs
}
