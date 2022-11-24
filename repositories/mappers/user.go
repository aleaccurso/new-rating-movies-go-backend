package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func UserModelToResDTO(model models.User) dtos.UserResDTO {
	return dtos.UserResDTO{
		Id:         model.Id,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		Nickname:   model.Nickname,
		Email:      model.Email,
		IsAdmin:    model.IsAdmin,
		Favorites:  model.Favorites,
		Rates:      RateModelsToResDtos(model.Rates),
		Language:   model.Language,
		ProfilePic: model.ProfilePic,
	}
}

func UserModelsToResDTOs(models []models.User) []dtos.UserResDTO {
	dtos := make([]dtos.UserResDTO, len(models))

	for i, model := range models {
		dtos[i] = UserModelToResDTO(model)
	}

	return dtos
}
