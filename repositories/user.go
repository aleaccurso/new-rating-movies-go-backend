package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database *mongo.Database
}

func InitialiseUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{
		database: db,
	}
}
