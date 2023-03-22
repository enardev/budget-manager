package usecase

import (
	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/errors"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port"
)

const ExpenseName = "expense"

type ExpenseUseCase struct {
	repository port.ExpenseRepository
	idGen      port.IdGenerator
}

func (uc ExpenseUseCase) FindByID(id string) (*model.Expense, error) {
	if !uc.repository.Exists(id) {
		return nil, errors.NewItemNotFoundError(ExpenseName)
	}
	return uc.repository.FindByID(id)
}

func (uc ExpenseUseCase) FindAll() ([]model.Expense, error) {
	return uc.repository.FindAll()
}

func (uc ExpenseUseCase) Save(expense model.Expense) (*model.Expense, error) {
	item := &expense
	item.Id = uc.idGen.GenerateID()

	result, err := uc.repository.Save(*item)
	if err != nil {
		return nil, errors.NewSaveItemError(ExpenseName)
	}

	return result, nil
}

func (uc ExpenseUseCase) Update(expense model.Expense) (*model.Expense, error) {
	if !uc.repository.Exists(expense.Id) {
		return nil, errors.NewItemNotFoundError(ExpenseName)
	}

	result, err := uc.repository.Update(expense)
	if err != nil {
		return nil, errors.NewUpdateItemError(ExpenseName)
	}

	return result, nil
}

func (uc ExpenseUseCase) Delete(id string) error {
	if !uc.repository.Exists(id) {
		return errors.NewItemNotFoundError(ExpenseName)
	}

	if err := uc.repository.Delete(id); err != nil {
		return errors.NewDeleteItemError(ExpenseName)
	}

	return nil
}
