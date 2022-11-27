package models

type MovieRate struct {
	Title string `bson:"title,omitempty" json:"title"`
	Key   string `bson:"key,omitempty" json:"key"`
}
