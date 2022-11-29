package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"new-rating-movies-go-backend/constants"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type AuthMiddleware struct{}

func InitialiseAuthMiddleware() IAuthMiddleware {
	return AuthMiddleware{}
}

func (middleware AuthMiddleware) Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gets token from header
		token := c.GetHeader("Authorization")

		// Verifies that the bearer token is included
		token, err := verifyToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}

		// Validates token
		secretKey := os.Getenv("JWT_SECRET")
		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, constants.AUTH_INVALID_TOKEN)
			return
		}

		// Verifies that the token is not nil
		if parsedToken == nil {
			c.IndentedJSON(http.StatusUnauthorized, constants.AUTH_INVALID_TOKEN)
			return
		}

		// Verifies claims
		roleClaims := claims["role"]
		emailClaims := claims["email"]

		roleString := fmt.Sprintf("%v", roleClaims)
		tokenRoles := strings.Split(roleString, ",")

		emailString := fmt.Sprintf("%v", emailClaims)

		if len(roles) > 0 {

			if roleClaims == nil {
				c.IndentedJSON(http.StatusForbidden, constants.AUTH_MISSING_PERMISSIONS)
				return
			}

			// Verifies that at least one role is in common
			if !lo.Some(tokenRoles, roles) {
				c.IndentedJSON(http.StatusForbidden, constants.AUTH_MISSING_PERMISSIONS)
				return
			}

			// Verifies that the account has been confirmed
			emailVerified := claims["email_verified"].(bool)
			if !emailVerified {
				c.IndentedJSON(http.StatusForbidden, constants.AUTH_UNVERIFIED_EMAIL)
				return
			}

		}

		// Save email and role in context
		c.Set("user_email", emailString)
		c.Set("user_role", roleString)

		// Calls the next handler in chain
		f(c)
	}
}

func verifyToken(token string) (resp string, error error) {
	// Verifies that a token has been given
	if token == "" {
		return "", errors.New(constants.AUTH_MISSING_TOKEN)
	}

	// Verifies that the token starts with bearer
	if !strings.Contains(strings.ToLower(token), "bearer") {
		return "", errors.New(constants.AUTH_WRONG_TOKEN_FORMAT)
	}

	// Returns the substring containing only the token
	return token[7:], nil
}
