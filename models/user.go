package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	Nickname  string        `json:"nickname"`
	Email     string        `json:"email"`
}
