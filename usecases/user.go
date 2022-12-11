package usecases

import (
	"context"
	"errors"
	"math"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"
	"new-rating-movies-go-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (usecase UserUsecase) GetUsers(c *gin.Context, page string, size string) (*dtos.UserPagingResDTO, error) {

	ctx := context.TODO()

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "page")
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "size")
	}

	usersCount, err := usecase.repository.UserRepository.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	nbPages := math.Ceil(float64(*usersCount)/float64(sizeInt))

	if nbPages == 0 {
		nbPages = 1
	}
	
	if float64(pageInt) >= nbPages - 1 {
		pageInt = int(nbPages - 1)
	}

	pagingUsers := dtos.UserPagingResDTO{
		Page:      int8(pageInt),
		Size:      int8(sizeInt),
		NbPages:   int8(nbPages),
		NbResults: int16(*usersCount),
	}

	users, err := usecase.repository.UserRepository.GetUsers(ctx, pageInt, sizeInt)
	if err != nil {
		return nil, err
	}

	pagingUsers.Data = mappers.UserModelsToResDTOs(users)

	return &pagingUsers, nil
}

func (usecase UserUsecase) GetUserById(c *gin.Context, userId string) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return nil, errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return nil, errors.New("cannot get logged user role")
	}

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}
	
	user, err := usecase.repository.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if user.Email != loggedUserEmail && loggedUserRole != "admin" {
		return nil, errors.New(constants.AUTH_UNAUTHORIZED)
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}

func (usecase UserUsecase) GetUserByEmail(c *gin.Context, email string) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return nil, errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return nil, errors.New("cannot get logged user role")
	}

	if !utils.IsEmailValid(email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	user, err := usecase.repository.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user.Email != loggedUserEmail && loggedUserRole != "admin" {
		return nil, errors.New(constants.AUTH_UNAUTHORIZED)
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}

func (usecase UserUsecase) UpdateUserById(c *gin.Context, userId string, reqUpdateDTO dtos.UserReqUpdateDTO) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return nil, errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return nil, errors.New("cannot get logged user role")
	}

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	userNewInfo := mappers.UserReqUpdateDTOToModel(reqUpdateDTO)

	existinguser, err := usecase.repository.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if existinguser.Email != loggedUserEmail && loggedUserRole != "admin" {
		return nil, errors.New(constants.AUTH_UNAUTHORIZED)
	}

	userNewInfo.Id = existinguser.Id
	userNewInfo.CreatedAt = existinguser.CreatedAt

	err = usecase.repository.UserRepository.ModifyUserById(ctx, userNewInfo)
	if err != nil {
		return nil, err
	}

	updatedUser, err := usecase.repository.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if reqUpdateDTO.Nickname != updatedUser.Nickname || reqUpdateDTO.Email != updatedUser.Email || reqUpdateDTO.Admin != updatedUser.IsAdmin || reqUpdateDTO.Language != updatedUser.Language || reqUpdateDTO.ProfilePic != updatedUser.ProfilePic {
		return nil, errors.New("something whent wrong during the update")
	}

	userResDTO := mappers.UserModelToResDTO(*updatedUser)

	return &userResDTO, nil
}

func (usecase UserUsecase) DeleteUserById(c *gin.Context, userId string) (*primitive.ObjectID, error) {

	ctx := context.TODO()

	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return nil, errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return nil, errors.New("cannot get logged user role")
	}

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	user, err := usecase.repository.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if user.Email != loggedUserEmail && loggedUserRole != "admin" {
		return nil, errors.New(constants.AUTH_UNAUTHORIZED)
	}

	err = usecase.repository.UserRepository.DeleteUserById(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return &user.Id, nil
}
