package repositories

import "gorm.io/gorm"

type Repository struct {
	database       *gorm.DB
	UserRepository UserRepository
}

func Initialise(database *gorm.DB) Repository {
	return Repository{
		database:       database,
		UserRepository: InitialiseUserRepository(database),
	}
}
