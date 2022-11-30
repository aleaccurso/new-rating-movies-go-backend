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

func (usecase UserUsecase) UpdateUserById(context context.Context, userId string, reqUpdateDTO dtos.UserReqUpdateDTO) (*dtos.UserResDTO, error) {

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	userNewInfo := mappers.UserReqUpdateDTOToModel(reqUpdateDTO)

	existinguser, err := usecase.repository.UserRepository.GetUserById(context, userUUID)
	if err != nil {
		return nil, err
	}

	userNewInfo.Id = existinguser.Id
	userNewInfo.CreatedAt = existinguser.CreatedAt
	userNewInfo.Favorites = existinguser.Favorites
	userNewInfo.Rates = existinguser.Rates

	err = usecase.repository.UserRepository.ModifyUserById(context, userNewInfo)
	if err != nil {
		return nil, err
	}

	updateduser, err := usecase.repository.UserRepository.GetUserById(context, userUUID)
	if err != nil {
		return nil, err
	}

	if reqUpdateDTO.Nickname != updateduser.Nickname || reqUpdateDTO.Email != updateduser.Email || reqUpdateDTO.Admin != updateduser.IsAdmin || reqUpdateDTO.Language != updateduser.Language || reqUpdateDTO.ProfilePic != updateduser.ProfilePic {
		// ! Password check is missing in the condition
		return nil, errors.New("something whent wrong during the update")
	}

	userResDTO := mappers.UserModelToResDTO(*updateduser)

	return &userResDTO, nil
}

func (usecase UserUsecase) DeleteUserById(context context.Context, userId string) (*primitive.ObjectID, error) {

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	user, err := usecase.repository.UserRepository.GetUserById(context, userUUID)
	if err != nil {
		return nil, err
	}

	err = usecase.repository.UserRepository.DeleteUserById(context, user.Id)
	if err != nil {
		return nil, err
	}

	return &user.Id, nil
}