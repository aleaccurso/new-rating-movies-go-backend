package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty"`
	CreatedAt primitive.Timestamp `bson:"createdAt,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt,omitempty"`
	Nickname  string              `bson:"nickename,omitempty"`
	Email     string              `bson:"email,omitempty"`
	Password  string              `bson:"password,omitempty"`
	IsAdmin   bool                `bson:"is_admin,omitempty"`
	Favorites []int32             `bson:"my_favorites,omitempty"`
	Rates     []Rate              `bson:"my_rates,omitempty"`
	Language  Language            `bson:"language,omitempty"`
	PicUrl    string              `bson:"profile_pic,omitempty"`
}
