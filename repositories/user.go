package repositories

import "gorm.io/gorm"

type UserRepository struct {
	database *gorm.DB
}

func InitialiseUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		database: db,
	}
}
