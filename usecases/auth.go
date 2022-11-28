package usecases

import (
	"context"
	"errors"
	"net/mail"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/enums"
	"new-rating-movies-go-backend/repositories"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (usecase AuthUsecase) Register(context context.Context, userDTO dtos.UserReqCreateDTO) (*primitive.ObjectID, error) {

	if !isEmailValid(userDTO.Email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	if !slices.Contains(enums.AllowedLanguages, userDTO.Language) {
		return nil, errors.New(constants.BAD_DATA + "language")
	}

	hashedPassword, err := getHash([]byte(userDTO.Password))
	if err != nil {
		return nil, err
	}
	userDTO.Password = *hashedPassword

	user, _ := usecase.repository.UserRepository.GetUserByEmail(context, userDTO.Email)
	if user != nil {
		return nil, errors.New(constants.AUTH_EMAIL_EXISTS)
	}

	newId, err := usecase.repository.AuthRepository.AddUser(context, userDTO)
	if err != nil {
		return nil, err
	}

	return newId, nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getHash(pwd []byte) (*string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil, errors.New(constants.AUTH_UNABLE_TO_HASH_PASSWORD)
	}
	return lo.ToPtr(string(hash)), nil
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
