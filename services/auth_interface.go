package services

import (
	"new-rating-movies-go-backend/dtos"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthService interface {
	Register(context *gin.Context, userDTO dtos.UserReqCreateDTO) (*primitive.ObjectID, error)
	Login(context *gin.Context, loginReqDTO dtos.LoginReqDTO) (*string, error)
	Logout(context *gin.Context)
}
