package dtos

type CheckMovieDBResDTO struct {
	MovieDbId int32 `bson:"movie_db_id" json:"movie_db_id"`
	IsInDB    bool  `bson:"is_in_db,omitempty" json:"is_in_db"`
}
