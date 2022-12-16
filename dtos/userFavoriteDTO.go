package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserFavoriteDTO struct {
	UserId    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	MovieDbId int32              `bson:"movie_db_id" json:"movie_db_id"`
}
