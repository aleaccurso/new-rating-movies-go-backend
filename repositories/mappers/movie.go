package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func ApiGetMovieInfoResDTOToMovieModel(dto dtos.ApiGetMovieInfoResDTO) models.Movie {
	return models.Movie{
		MovieDbId:   dto.MovieDbId,
		ReleaseDate: dto.ReleaseDate,
		Director:    dto.Director,
		Casting:     dto.Casting,
		VoteAverage: int8(dto.VoteAverage),
		VoteCount:   int32(dto.VoteCount),
		Genre:       dto.Genre,
		En:          LocalMovieInfoResDTOToModel(dto.En),
		Fr:          LocalMovieInfoResDTOToModel(dto.Fr),
		It:          LocalMovieInfoResDTOToModel(dto.It),
		Nl:          LocalMovieInfoResDTOToModel(dto.Nl),
	}
}

func LocalMovieInfoResDTOToModel(dto dtos.LocalMovieInfoResDTO) models.MovieLocalInfo {
	return models.MovieLocalInfo{
		Title:      dto.Title,
		Overview:   dto.Overview,
		PosterPath: dto.PosterPath,
		Trailers:   LocalMovieTrailerResDTOsToModels(dto.Trailers),
	}
}

func LocalMovieTrailerResDTOsToModels(trailerDTOs []dtos.LocalMovieTrailerResDTO) []models.MovieTrailer {
	trailers := make([]models.MovieTrailer, len(trailerDTOs))

	for i, trailerDTO := range trailerDTOs {
		trailers[i] = LocalMovieTrailerResDTOToModel(trailerDTO)
	}

	return trailers
}

func LocalMovieTrailerResDTOToModel(trailerDTO dtos.LocalMovieTrailerResDTO) models.MovieTrailer {
	return models.MovieTrailer{
		Title: trailerDTO.Title,
		Key:   trailerDTO.Key,
		Site:  trailerDTO.Site,
		Type:  trailerDTO.Type,
	}
}
