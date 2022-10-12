package middlewares

import "github.com/gin-gonic/gin"

type AuthMiddleware struct{}

type IAuthMiddleware interface {
	Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc
}

func InitialiseAuthMiddleware() IAuthMiddleware {
	return AuthMiddleware{}
}

func (middleware AuthMiddleware) Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gets token from header
		// token := c.GetHeader("Authorization")

		// Calls the next handler in chain
		f(c)
	}
}
