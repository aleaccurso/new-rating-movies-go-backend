package repositories

import (
	"new-rating-movies-go-backend/database"
)

type Repository struct {
	RepositoryBase
}

func Initialise(database *database.Database) Repository {
	return Repository{
		RepositoryBase: RepositoryBase{
			UserRepository: InitialiseUserRepository(database),
			AuthRepository: InitialiseAuthRepository(database),
		},
	}
}
