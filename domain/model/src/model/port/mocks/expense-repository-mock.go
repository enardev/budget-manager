package mocks

import "github.com/enaldo1709/budget-manager/domain/model/src/model"

type ExpenseRepositoryMock struct {
	ExistsFn   func(string) bool
	FindByIDFn func(string) (*model.Expense, error)
	FindAllFn  func() ([]model.Expense, error)
	SaveFn     func(model.Expense) (*model.Expense, error)
	UpdateFn   func(model.Expense) (*model.Expense, error)
	DeleteFn   func(string) error
}

func (m *ExpenseRepositoryMock) Exists(id string) bool {
	return m.ExistsFn(id)
}

func (m *ExpenseRepositoryMock) FindByID(id string) (*model.Expense, error) {
	return m.FindByIDFn(id)
}

func (m *ExpenseRepositoryMock) FindAll() ([]model.Expense, error) {
	return m.FindAllFn()
}

func (m *ExpenseRepositoryMock) Save(e model.Expense) (*model.Expense, error) {
	return m.SaveFn(e)
}

func (m *ExpenseRepositoryMock) Update(e model.Expense) (*model.Expense, error) {
	return m.UpdateFn(e)
}

func (m *ExpenseRepositoryMock) Delete(id string) error {
	return m.DeleteFn(id)
}
