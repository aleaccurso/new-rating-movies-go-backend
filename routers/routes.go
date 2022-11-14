package routers

import (
	"new-rating-movies-go-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine     *gin.Engine
	controller controllers.Controller
}

func Initialise(engine *gin.Engine, controller controllers.Controller) Router {
	return Router{
		engine:     engine,
		controller: controller,
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

	api.GET("/users", router.controller.UserController.GetUsers)
	api.GET("/users/:userId", router.controller.UserController.GetUserById)

	// Runs the engine
	if err := router.engine.Run(":8010"); err != nil {
		return err
	}

	return nil
}
