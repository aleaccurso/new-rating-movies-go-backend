package dtos

type UserResDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"modified_at"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
}
