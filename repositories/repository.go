package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database       *mongo.Database
	UserRepository UserRepository
}

func Initialise(database *mongo.Database) Repository {
	return Repository{
		database:       database,
		UserRepository: InitialiseUserRepository(database),
	}
}
