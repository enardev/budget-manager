package port

import (
	"github.com/enaldo1709/budget-manager/domain/model/src/model"
)

type ExpenseRepository interface {
	Exists(id int) bool
	FindByID(id int) (*model.Expense, error)
	FindAll() ([]model.Expense, error)
	Save(*model.Expense) (*model.Expense, error)
	Update(*model.Expense) (*model.Expense, error)
	Delete(id int) error
}
