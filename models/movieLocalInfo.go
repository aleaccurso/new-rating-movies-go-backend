package models

type MovieLocalInfo struct {
	Title      string      `bson:"title,omitempty" json:"title"`
	Overview   string      `bson:"overview,omitempty" json:"overview"`
	PosterPath string      `bson:"poster_path,omitempty" json:"poster_path"`
	Trailers   []MovieRate `bson:"trailers,omitempty" json:"trailers"`
}
