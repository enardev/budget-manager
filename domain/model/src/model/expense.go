package model

import (
	"time"
)

type Expense struct {
	Id     string    `json:"id" validate:"required,uuid4"`
	Amount float64   `json:"amount" validate:"required,number"`
	Date   time.Time `json:"date" validate:"required"`
}
