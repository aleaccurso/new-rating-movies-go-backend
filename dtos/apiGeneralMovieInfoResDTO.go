package dtos

type ApiGeneralMovieInfoResDTO struct {
	Id          int32         `bson:"id,omitempty" json:"id"`
	Genres      []ApiGenreDTO `bson:"genres,omitempty" json:"genres"`
	VoteAverage float32       `bson:"vote_average,omitempty" json:"vote_average"`
	VoteCount   float32       `bson:"vote_count,omitempty" json:"vote_count"`
	ReleaseDate string        `bson:"release_date,omitempty" json:"release_date"`
	Credits     ApiCreditsDTO `bson:"credits,omitempty" json:"credits"`
}

type ApiGenreDTO struct {
	Name string `bson:"name,omitempty" json:"name"`
}

type ApiCreditsDTO struct {
	Crew []CrewDTO `bson:"crew,omitempty" json:"crew"`
	Cast []CastDTO `bson:"cast,omitempty" json:"cast"`
}
