package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func RateModelToResDto(model models.UserRate) dtos.RateResDTO {
	return dtos.RateResDTO{
		MovieDbId: model.MovieDbId,
		Rate:      model.UserRate,
	}
}

func RateModelsToResDtos(models []models.UserRate) []dtos.RateResDTO {
	dtos := make([]dtos.RateResDTO, len(models))

	for i, model := range models {
		dtos[i] = RateModelToResDto(model)
	}

	return dtos
}
