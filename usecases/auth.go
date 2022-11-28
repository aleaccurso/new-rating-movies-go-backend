package usecases

import (
	"context"
	"errors"
	"log"
	"net/mail"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/enums"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

type AuthUsecase struct {
	repository repositories.Repository
}

func InitialiseAuthUsecase(repository repositories.Repository) AuthUsecase {
	return AuthUsecase{
		repository: repository,
	}
}

func (usecase AuthUsecase) Register(context context.Context, userDTO dtos.UserReqCreateDTO) (*dtos.UserResDTO, error) {

	if !isEmailValid(userDTO.Email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	if !slices.Contains(enums.AllowedLanguages, userDTO.Language) {
		return nil, errors.New(constants.BAD_DATA + "language")
	}

	userDTO.Password = getHash([]byte(userDTO.Password))

	user, err := usecase.repository.AuthRepository.AddUser(context, userDTO)
	if err != nil {
		return nil, err
	}

	insertedUserDTO := mappers.UserModelToResDTO(*user)

	return &insertedUserDTO, nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// func generateJWT() (string, error) {
// 	secretKey := os.Getenv("JWT_SECRET")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		log.Println("Error in JWT token generation")
// 		return "", err
// 	}
// 	return tokenString, nil
// }
