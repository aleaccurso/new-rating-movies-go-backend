package dtos

type ApiMovieVideoResDTO struct {
	Id      int32                     `bson:"id,omitempty" json:"id"`
	Results []LocalMovieTrailerResDTO `bson:"results,omitempty" json:"results"`
}
