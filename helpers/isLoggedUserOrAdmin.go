package helpers

import (
	"errors"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsLoggedUserOrAdmin(c *gin.Context, userId primitive.ObjectID, user models.User) error {
	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return errors.New("cannot get logged user role")
	}

	if user.Email != loggedUserEmail && loggedUserRole != "admin" {
		return errors.New(constants.AUTH_UNAUTHORIZED)
	}

	return nil
}
