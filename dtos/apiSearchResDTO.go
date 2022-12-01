package dtos

type ApiSearchResDTO struct {
	Page         int16               `bson:"page,omitempty" json:"page"`
	Results      []ApiSearchMovieDTO `bson:"results,omitempty" json:"results"`
	TotalPages   int16               `bson:"total_pages,omitempty" json:"total_pages"`
	TotalResults int32               `bson:"total_results,omitempty" json:"total_results"`
}
