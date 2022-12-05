package dtos

type ApiLocalMovieTrailerResDTO struct {
	Name string `bson:"name,omitempty" json:"name"`
	Key  string `bson:"key,omitempty" json:"key"`
	Site string `bson:"site,omitempty" json:"site"`
	Type string `bson:"type,omitempty" json:"type"`
}
