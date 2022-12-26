package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func rateModelToResDto(model models.UserRate) dtos.RateResDTO {
	return dtos.RateResDTO{
		MovieDbId: model.MovieDbId,
		Rate:      model.UserRate,
	}
}

func rateModelsToResDtos(models []models.UserRate) []dtos.RateResDTO {
	dtos := make([]dtos.RateResDTO, len(models))

	for i, model := range models {
		dtos[i] = rateModelToResDto(model)
	}

	return dtos
}

func RateReqUpdateDTOToModel(dto dtos.UserRateReqUpdateDTO) models.UserRate {
	return models.UserRate{
		MovieDbId: dto.MovieDbId,
		UserRate:  dto.Rate,
	}
}

func RateReqUpdateDTOToRateResDTO(reqDTO dtos.UserRateReqUpdateDTO) dtos.RateResDTO {
	return dtos.RateResDTO{
		MovieDbId: reqDTO.MovieDbId,
		Rate:  reqDTO.Rate,
	}
}