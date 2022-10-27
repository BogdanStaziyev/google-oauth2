package domain

type User struct {
	GoogleID string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	//Password string `json:"password" validate:"required"`
}
