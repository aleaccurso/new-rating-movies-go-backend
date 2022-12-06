package usecases

import (
	"new-rating-movies-go-backend/dtos"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUsecase interface {
	GetUsers(c *gin.Context, page string, size string) (*dtos.UserPagingResDTO, error)
	GetUserById(c *gin.Context, userId string) (*dtos.UserResDTO, error)
	GetUserByEmail(c *gin.Context, email string) (*dtos.UserResDTO, error)
	UpdateUserById(c *gin.Context, userId string, reqUpdateDTO dtos.UserReqUpdateDTO) (*dtos.UserResDTO, error)
	DeleteUserById(c *gin.Context, userId string) (*primitive.ObjectID, error)
}
