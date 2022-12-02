package dtos

type ApiGetMovieInfoResDTO struct {
	MovieDbId   int32    `bson:"movie_db_id,omitempty" json:"movie_db_id"`
	ReleaseDate string   `bson:"release_date,omitempty" json:"release_date"`
	Director    string   `bson:"director,omitempty" json:"director"`
	Casting     string   `bson:"casting,omitempty" json:"casting"`
	VoteAverage float32  `bson:"vote_average,omitempty" json:"vote_average"`
	VoteCount   float32  `bson:"vote_count,omitempty" json:"vote_count"`
	Genre       []string `bson:"genre,omitempty" json:"genre"`

	En ApiLocalMovieInfoResDTO `bson:"en,omitempty" json:"en"`
	Fr ApiLocalMovieInfoResDTO `bson:"fr,omitempty" json:"fr"`
	It ApiLocalMovieInfoResDTO `bson:"it,omitempty" json:"it"`
	Nl ApiLocalMovieInfoResDTO `bson:"nl,omitempty" json:"nl"`
}
