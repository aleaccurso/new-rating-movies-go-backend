package mappers

import (
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

func UserReqCreateDTOToModel(dto dtos.UserReqCreateDTO) models.User {
	return models.User{
		Nickname:   dto.Nickname,
		Email:      dto.Email,
		Password:   dto.Password,
		IsAdmin:    false,
		Favorites:  nil,
		Rates:      nil,
		Language:   dto.Language,
		ProfilePic: "",
	}
}

func UserReqUpdateDTOToModel(dto dtos.UserReqUpdateDTO) models.User {
	return models.User{
		Nickname:   dto.Nickname,
		Email:      dto.Email,
		Password:   dto.Password,
		IsAdmin:    dto.Admin,
		Language:   dto.Language,
		ProfilePic: dto.ProfilePic,
	}
}

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
