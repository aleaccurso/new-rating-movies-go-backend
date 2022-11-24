package dtos

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserResDTO struct {
	Id         bson.ObjectId `json:"id"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"modified_at"`
	Nickname   string        `json:"nickname"`
	Email      string        `json:"email"`
	IsAdmin    bool          `bson:"is_admin,omitempty" json:"is_admin"`
	Favorites  []int32       `bson:"my_favorites,omitempty" json:"my_favorites"`
	Rates      []RateResDTO  `bson:"my_rates,omitempty" json:"my_rates"`
	Language   string        `bson:"language,omitempty" json:"language"`
	ProfilePic string        `bson:"profile_pic,omitempty" json:"profile_pic"`
}
