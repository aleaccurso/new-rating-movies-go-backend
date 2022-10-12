package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Used by router to call the middlewares
func UseMiddlewares(api *gin.RouterGroup) gin.IRoutes {
	// sortMiddlewares := make([]SortMiddleware, 0, len(MiddlewaresBase)+len(Middlewares))
	// sortMiddlewares = append(sortMiddlewares, Middlewares...)

	// // Comment this to disable base middlewares 🔧
	// sortMiddlewares = append(sortMiddlewares, MiddlewaresBase...)

	// // Sort the middlewares
	// sort.Slice(sortMiddlewares, func(i, j int) bool {
	// 	return sortMiddlewares[i].Order < sortMiddlewares[j].Order
	// })

	// // Get the handlers
	// middlewares := lo.Map(sortMiddlewares, func(middleware SortMiddleware, _ int) gin.HandlerFunc {
	// 	return middleware.handler
	// })

	// // Use the middlewares
	return api.Use()
}

// List of custom middlewares
// var Middlewares = []SortMiddleware{}
