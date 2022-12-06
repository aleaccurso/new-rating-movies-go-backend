package models

type MovieTrailer struct {
	Title string `bson:"title,omitempty" json:"title"`
	Key   string `bson:"key,omitempty" json:"key"`
	Site  string `bson:"site,omitempty" json:"site"`
	Type  string `bson:"type,omitempty" json:"type"`
}
