package dtos

type ApiMovieVideoResDTO struct {
	Id      int32                        `bson:"id,omitempty" json:"id"`
	Results []ApiLocalMovieTrailerResDTO `bson:"results,omitempty" json:"results"`
}
