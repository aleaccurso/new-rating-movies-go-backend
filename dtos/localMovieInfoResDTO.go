package dtos

type LocalMovieInfoResDTO struct {
	PosterPath string                    `bson:"poster_path,omitempty" json:"poster_path"`
	Title      string                    `bson:"title,omitempty" json:"title"`
	Overview   string                    `bson:"overview,omitempty" json:"overview"`
	Trailers   []LocalMovieTrailerResDTO `bson:"trailers,omitempty" json:"trailers"`
}
