package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResDTO struct {
	Id         primitive.ObjectID `bson:"id,omitempty" json:"id"`
	CreatedAt  time.Time          `bson:"created_a,omitempty" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	Nickname   string             `bson:"nickname,omitempty" json:"nickname"`
	Email      string             `bson:"email,omitempty" json:"email"`
	IsAdmin    bool               `bson:"is_admin,omitempty" json:"is_admin"`
	Favorites  []int32            `bson:"my_favorites,omitempty" json:"my_favorites"`
	Rates      []RateResDTO       `bson:"my_rates,omitempty" json:"my_rates"`
	Language   string             `bson:"language,omitempty" json:"language"`
	ProfilePic string             `bson:"profile_pic,omitempty" json:"profile_pic"`
}
