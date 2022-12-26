package dtos

type UserRateReqUpdateDTO struct {
	MovieDbId int32 `bson:"movie_db_id" json:"movie_db_id"`
	Rate      float32  `bson:"rate" json:"rate"`
}