package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"strings"
)

func ApiMovieInfoToMovieResDTO(general dtos.ApiGeneralMovieInfoResDTO, local map[string]dtos.LocalMovieInfoResDTO) (*dtos.ApiGetMovieInfoResDTO, error) {

	director := ""

	for _, crew := range general.Credits.Crew {
		if crew.Job == "Director" {
			director = crew.Name
			break
		}
	}

	casting := toCastingString(general.Credits.Cast)

	genres := []string{}

	for _, genre := range general.Genres {
		if genre.Name != "" {
			genres = append(genres, genre.Name)
		}
	}

	return &dtos.ApiGetMovieInfoResDTO{
		MovieDbId:   general.Id,
		ReleaseDate: general.ReleaseDate,
		Director:    director,
		Casting:     casting,
		VoteAverage: general.VoteAverage,
		VoteCount:   general.VoteCount,
		Genre:       genres,
		En:          local["en"],
		Fr:          local["fr"],
		It:          local["it"],
		Nl:          local["nl"],
	}, nil
}

func toCastingString(casting []dtos.CastDTO) string {

	count := 0
	topThreeActors := []string{}

	for _, cast := range casting {
		if cast.Name != "" {
			topThreeActors = append(topThreeActors, cast.Name)
			count++
		}
		if count == 3 {
			break
		}
	}

	return strings.Join(topThreeActors, " - ")
}
