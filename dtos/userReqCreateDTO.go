package dtos

type UserReqCreateDTO struct {
	Nickname string `bson:"nickname,omitempty" json:"nickname"`
	Email    string `bson:"email,omitempty" json:"email"`
	Password string `bson:"password,omitempty" json:"password"`
	Language string `bson:"language,omitempty" json:"language"`
}
