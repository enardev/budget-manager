package port

import (
	"github.com/enaldo1709/budget-manager/domain/model/src/model"
)

type ExpenseRepository interface {
	Exists(id string) bool
	FindByID(id string) (*model.Expense, error)
	FindAll() ([]model.Expense, error)
	Save(model.Expense) (*model.Expense, error)
	Update(model.Expense) (*model.Expense, error)
	Delete(id string) error
}
