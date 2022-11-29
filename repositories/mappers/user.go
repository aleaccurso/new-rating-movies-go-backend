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
