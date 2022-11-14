package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
	Nickname    string        `json:"nickname"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	IsAdmin     bool          `json:"is_admin"`
	Favorites []int32       `json:"my_favorites"`
	Rates     []Rate        `json:"my_rates"`
	Language    string        `json:"language"`
	ProfilePic  string        `json:"profile_pic"`
}
