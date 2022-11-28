package dtos

type RateResDTO struct {
	MovieDbId int32 `bson:"movie_db_id,omitempty" json:"movie_db_id"`
	Rate      int8  `bson:"rate,omitempty" json:"rate"`
}
