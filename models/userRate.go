package models

type UserRate struct {
	MovieDbId int32 `bson:"movie_db_id,omitempty" json:"movie_db_id"`
	UserRate  int8  `bson:"rate,omitempty" json:"rate"`
}
