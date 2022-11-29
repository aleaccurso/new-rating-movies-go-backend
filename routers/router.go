package routers

import (
	"new-rating-movies-go-backend/controllers"
	middlewares "new-rating-movies-go-backend/middelwares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine         *gin.Engine
	authMiddleware middlewares.IAuthMiddleware
	controller     controllers.Controller
}

func Initialise(engine *gin.Engine, authMiddleware middlewares.IAuthMiddleware, controller controllers.Controller) Router {
	return Router{
		engine:         engine,
		authMiddleware: authMiddleware,
		controller:     controller,
	}
}

func (router Router) Run() error {

	// Setup CORS
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Creates the api-group
	api := router.engine.Group("/api")

	////////////////////////////////////////
	//     Initialises all the routers    //
	////////////////////////////////////////

	// Authentication
	api.GET("/me", router.authMiddleware.Authorize(router.controller.AuthController.GetMe))
	api.GET("/login", router.controller.AuthController.Login)
	api.GET("/logout",  router.authMiddleware.Authorize(router.controller.AuthController.Logout))
	api.POST("/register", router.controller.AuthController.Register)

	// User
	api.GET("/users", router.controller.UserController.GetUsers)
	api.GET("/users/:userId", router.controller.UserController.GetUserById)
	api.PATCH("/users/:userId", router.controller.UserController.GetUserById)
	api.DELETE("/users/:userId", router.controller.UserController.GetUserById)

	// Movie
	api.GET("/movies", router.controller.UserController.GetUsers)
	api.POST("/movies", router.controller.UserController.GetUsers)
	api.GET("/movies/:movieId", router.controller.UserController.GetUserById)
	api.PATCH("/movies/:movieId", router.controller.UserController.GetUserById)
	api.DELETE("/movies/:movieId", router.controller.UserController.GetUserById)

	// Run the engine
	if err := router.engine.Run(":8010"); err != nil {
		return err
	}

	return nil
}
