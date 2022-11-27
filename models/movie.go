package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	Id          bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	MovieDbId   int32          `bson:"movie_db_id,omitempty" json:"movie_db_id"`
	ReleaseDate string         `bson:"release_date,omitempty" json:"release_date"`
	Director    string         `bson:"director,omitempty" json:"director"`
	Casting     string         `bson:"casting,omitempty" json:"casting"`
	VoteAverage int8           `bson:"vote_average,omitempty" json:"vote_average"`
	VoteCount   int32          `bson:"vote_count,omitempty" json:"vote_count"`
	Genre       []string       `bson:"genre,omitempty" json:"genre"`
	En          MovieLocalInfo `bson:"en,omitempty" json:"en"`
	Fr          MovieLocalInfo `bson:"fr,omitempty" json:"fr"`
	It          MovieLocalInfo `bson:"it,omitempty" json:"it"`
	Nl          MovieLocalInfo `bson:"nl,omitempty" json:"nl"`
}
