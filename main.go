package main

import (
	"fmt"
	"new-rating-movies-go-backend/controllers"
	db "new-rating-movies-go-backend/database"
	middlewares "new-rating-movies-go-backend/middelwares"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/routers"
	"new-rating-movies-go-backend/services"
	"new-rating-movies-go-backend/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	// Initialises the router
	engine := gin.Default()

	// Connects to the database
	database, err := db.Initialise()
	if err != nil {
		return fmt.Errorf("router: %s", err)
	}

	// Creates the repository container
	repository := repositories.Initialise(database)

	// Creates the service container
	service := services.Initialise(repository)

	// Creates the usecase container
	usecase := usecases.Initialise(repository, service)

	// Creates the controller container
	controller := controllers.Initialise(usecase, service)

	// Creates middlewares
	middleware := middlewares.Initialise()

	// Creates the routes container
	router := routers.Initialise(engine, middleware.AuthMiddleware, controller)

	// Start the router
	if err := router.Run(); err != nil {
		return fmt.Errorf("router: %s", err)
	}

	// defer db.CloseConnection(database.Client)

	return nil
}
