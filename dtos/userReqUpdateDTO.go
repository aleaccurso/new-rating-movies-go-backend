package dtos

type UserReqUpdateDTO struct {
	Nickname   string `bson:"nickname,omitempty" json:"nickname"`
	Email      string `bson:"email,omitempty" json:"email"`
	Admin      bool   `bson:"is_admin,omitempty" json:"is_admin"`
	Password   string `bson:"password,omitempty" json:"password"`
	Language   string `bson:"language,omitempty" json:"language"`
	ProfilePic string `bson:"profile_pic,omitempty" json:"profile_pic"`
}
