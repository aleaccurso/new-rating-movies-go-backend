package dtos

type RateResDTO struct {
	MovieDbId int32 `bson:"movie_db_id" json:"movie_db_id"`
	Rate      float32  `bson:"rate,omitempty" json:"rate"`
}
