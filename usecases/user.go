package usecases

import (
	"context"
	"errors"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"
	"new-rating-movies-go-backend/utils"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	repository repositories.Repository
}

func InitialiseUserUsecase(repository repositories.Repository) UserUsecase {
	return UserUsecase{
		repository: repository,
	}
}

func (usecase UserUsecase) GetUsers(context context.Context, page string, size string) ([]dtos.UserResDTO, error) {

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "page")
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "size")
	}

	users, err := usecase.repository.UserRepository.GetUsers(context, pageInt, sizeInt)
	if err != nil {
		return nil, err
	}

	return mappers.UserModelsToResDTOs(users), nil
}

func (usecase UserUsecase) GetUserById(context context.Context, userId string) (*dtos.UserResDTO, error) {

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}
	user, err := usecase.repository.UserRepository.GetUserById(context, userUUID)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}

func (usecase UserUsecase) GetUserByEmail(context context.Context, email string) (*dtos.UserResDTO, error) {

	if !utils.IsEmailValid(email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	user, err := usecase.repository.UserRepository.GetUserByEmail(context, email)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}