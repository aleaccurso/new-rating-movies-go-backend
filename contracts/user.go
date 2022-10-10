package contracts

type UserResDTO struct {
	ID         string  `json:"id"`
	CreatedAt  string  `json:"created_at"`
	ModifiedAt string  `json:"modified_at"`
	Nickname   string  `json:"nickname"`
	Email     string  `json:"email"`
	Price      float64 `json:"price"`
}
