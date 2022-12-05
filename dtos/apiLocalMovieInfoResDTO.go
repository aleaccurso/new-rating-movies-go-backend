package dtos

type ApiLocalMovieInfoResDTO struct {
	PosterPath string                       `bson:"poster_path,omitempty" json:"poster_path"`
	Title      string                       `bson:"title,omitempty" json:"title"`
	Overview   string                       `bson:"overview,omitempty" json:"overview"`
	Trailers   []ApiLocalMovieTrailerResDTO `bson:"trailers,omitempty" json:"trailers"`
}
