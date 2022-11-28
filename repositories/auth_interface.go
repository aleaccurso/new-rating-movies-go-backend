package repositories

import (
	"context"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/models"
)

type IAuthRepository interface {
	AddUser(context context.Context, user dtos.UserReqCreateDTO) (*models.User, error)
}
