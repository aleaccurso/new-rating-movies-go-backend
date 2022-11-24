package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedAt  time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
	Nickname   string        `bson:"nickname,omitempty" json:"nickname"`
	Email      string        `bson:"email,omitempty" json:"email"`
	Password   string        `bson:"password,omitempty" json:"password"`
	IsAdmin    bool          `bson:"is_admin,omitempty" json:"is_admin"`
	Favorites  []int32       `bson:"my_favorites,omitempty" json:"my_favorites"`
	Rates      []Rate        `bson:"my_rates,omitempty" json:"my_rates"`
	Language   string        `bson:"language,omitempty" json:"language"`
	ProfilePic string        `bson:"profile_pic,omitempty" json:"profile_pic"`
}
