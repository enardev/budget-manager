package usecase

import (
	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/errors"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port"
)

const ExpenseName = "expense"

type ExpenseUseCase struct {
	Repository port.ExpenseRepository
}

func (uc ExpenseUseCase) FindByID(id int) (*model.Expense, error) {
	if !uc.Repository.Exists(id) {
		return nil, errors.NewItemNotFoundError(ExpenseName)
	}
	return uc.Repository.FindByID(id)
}

func (uc ExpenseUseCase) FindAll() ([]model.Expense, error) {
	return uc.Repository.FindAll()
}

func (uc ExpenseUseCase) Save(expense *model.Expense) (*model.Expense, error) {
	if expense.Id < 0 {
		return nil, errors.NewInvalidItemError(ExpenseName, "field Id must be a positive integer")
	}

	if uc.Repository.Exists(expense.Id) {
		return nil, errors.NewItemAlreadyExistsError(ExpenseName)
	}

	result, err := uc.Repository.Save(expense)
	if err != nil {
		return nil, errors.NewSaveItemError(ExpenseName)
	}

	return result, nil
}

func (uc ExpenseUseCase) Update(expense *model.Expense) (*model.Expense, error) {
	if !uc.Repository.Exists(expense.Id) {
		return nil, errors.NewItemNotFoundError(ExpenseName)
	}

	result, err := uc.Repository.Update(expense)
	if err != nil {
		return nil, errors.NewUpdateItemError(ExpenseName)
	}

	return result, nil
}

func (uc ExpenseUseCase) Delete(id int) error {
	if !uc.Repository.Exists(id) {
		return errors.NewItemNotFoundError(ExpenseName)
	}

	if err := uc.Repository.Delete(id); err != nil {
		return errors.NewDeleteItemError(ExpenseName)
	}

	return nil
}
