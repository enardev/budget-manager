package model

import (
	"time"
)

type Expense struct {
	Id      int       `json:"id" validate:"integer"`
	Amount  float64   `json:"amount" validate:"required,number"`
	Created time.Time `json:"created" validate:"required"`
}
