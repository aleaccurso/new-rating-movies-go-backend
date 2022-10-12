package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Rate struct {
	MovieDbId int32 `json:"movieDbId"`
	Rate      int8  `json:"rate"`
}

type User struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	ModifiedAt  time.Time     `json:"modified_at"`
	Nickname    string        `json:"nickname"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	IsAdmin     bool          `json:"is_admin"`
	MyFavorites []int32       `json:"my_favorites"`
	MyRates     []Rate        `json:"my_rates"`
	Language    string        `json:"language"`
	ProfilePic  string        `json:"profile_pic"`
}
