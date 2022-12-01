package dtos

type ApiSearchMovieDTO struct {
	Id          int32  `bson:"id,omitempty" json:"id"`
	ReleaseDate string `bson:"release_date,omitempty" json:"release_date"`
	PosterPath  string `bson:"poster_path,omitempty" json:"poster_path"`
	Title       string `bson:"title,omitempty" json:"title"`
	Overview    string `bson:"overview,omitempty" json:"overview"`
}
