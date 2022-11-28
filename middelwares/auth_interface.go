package middlewares

import (
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc
}
