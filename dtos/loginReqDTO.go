package dtos

type LoginReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
