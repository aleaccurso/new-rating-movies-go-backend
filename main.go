package main

import (
<<<<<<< HEAD
	"fmt"
	"new-rating-movies-go-backend/controllers"
	"new-rating-movies-go-backend/database"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/routers"
	"new-rating-movies-go-backend/usecases"
=======
	"new-rating-movies-go-backend/controllers"
>>>>>>> 93ddb8e (Postman request ok - GetUsers - GetUserById)

	"github.com/gin-gonic/gin"
)

func main() {
<<<<<<< HEAD
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	// Initialises the router
	engine := gin.Default()

	// Connects to the database
	database := database.Initialise()

	// Creates the repository container
	repository := repositories.Initialise(database)

	// Creates the usecase container
	usecase := usecases.Initialise(repository)

	// Creates the controller container
	controller := controllers.Initialise(usecase)

	// Creates the routes container
	router := routers.Initialise(engine, controller)

	// Start the router
	if err := router.Run(); err != nil {
		return fmt.Errorf("router: %s", err)
	}

	return nil
=======
	router := gin.Default()

	router.GET("/api/users", controllers.GetUsers)
	router.GET("/api/users/:id", controllers.GetUserById)

	router.Run("localhost:8080")
>>>>>>> 93ddb8e (Postman request ok - GetUsers - GetUserById)
}
