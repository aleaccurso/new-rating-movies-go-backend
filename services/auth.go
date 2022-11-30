package services

import (
	"errors"
	"log"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"
	"new-rating-movies-go-backend/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository repositories.Repository
}

type JWTClaim struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	EmailVerified bool   `json:"email_verified"`
	jwt.StandardClaims
}

func InitialiseAuthService(repository repositories.Repository) AuthService {
	return AuthService{
		repository: repository,
	}
}

func (service AuthService) Register(context *gin.Context, userDTO dtos.UserReqCreateDTO) (*primitive.ObjectID, error) {

	user := mappers.UserReqCreateDTOToModel(userDTO)

	if !utils.IsEmailValid(userDTO.Email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	if !utils.IsAllowedLanguage(userDTO.Language) {
		return nil, errors.New(constants.BAD_DATA + "language")
	}

	hashedPassword, err := service.getHash([]byte(userDTO.Password))
	if err != nil {
		return nil, err
	}
	user.Password = *hashedPassword

	checkUser, _ := service.repository.UserRepository.GetUserByEmail(context, userDTO.Email)
	if checkUser != nil {
		return nil, errors.New(constants.AUTH_EMAIL_EXISTS)
	}

	newId, err := service.repository.AuthRepository.AddUser(context, user)
	if err != nil {
		return nil, err
	}

	return newId, nil
}

func (service AuthService) Login(context *gin.Context, loginReqDTO dtos.LoginReqDTO) (*string, error) {

	// check if email exists and password is correct
	user, err := service.repository.UserRepository.GetUserByEmail(context, loginReqDTO.Email)
	if err != nil {
		// c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return nil, err
	}

	credentialError := service.checkPassword(user.Password, loginReqDTO.Password)
	if credentialError != nil {
		// context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		// context.Abort()
		// return

		return nil, errors.New(constants.AUTH_PASSWORD_MISSMATCH)
	}

	tokenString, err := service.generateJWT(user.Nickname, user.Email, user.IsAdmin)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// context.Abort()
		return nil, err
	}

	return &tokenString, nil
}

func (service AuthService) Logout(context *gin.Context) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(context.Writer, &cookie)
}

func (service AuthService) getHash(pwd []byte) (*string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		return nil, errors.New(constants.AUTH_UNABLE_TO_HASH_PASSWORD)
	}
	return lo.ToPtr(string(hash)), nil
}

func (service AuthService) checkPassword(userPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (service AuthService) generateJWT(nickname string, email string, isAdmin bool) (string, error) {

	secretKey := os.Getenv("JWT_SECRET")

	expirationTime := time.Now().Add(24 * time.Hour)

	role := "user"

	if isAdmin {
		role = "admin"
	}

	claims := &JWTClaim{
		Email:         email,
		Username:      nickname,
		Role:          role,
		EmailVerified: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}

	return tokenString, nil
}
