package dto

type ExpenseDTO struct {
	Id      int     `json:"id" validate:"integer"`
	Amount  float64 `json:"amount" validate:"required,number"`
	Created string  `json:"created" validate:"required"`
}
