package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port/mocks"
)

func TestExpenseUseCaseFindByID(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Expense
		wantErr bool
	}{
		{
			name: "given an id then get a expense model",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(s int) bool {
						return true
					},
					FindByIDFn: func(id int) (*model.Expense, error) {
						return &model.Expense{
							Id:      1,
							Amount:  25.3,
							Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
						}, nil
					},
				},
			},
			args: args{id: 1},
			want: &model.Expense{
				Id:      1,
				Amount:  25.3,
				Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "given an id when the expense not exists then get an error",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(s int) bool {
						return false
					},
				},
			},
			args:    args{id: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				Repository: tt.fields.repository,
			}
			got, err := uc.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpenseUseCase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpenseUseCase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpenseUseCaseFindAll(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Expense
		wantErr bool
	}{
		{
			name: "Got an array of expenses",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					FindAllFn: func() ([]model.Expense, error) {
						return []model.Expense{
							{
								Id:      1,
								Amount:  33.5,
								Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
							},
							{
								Id:      1,
								Amount:  24.7,
								Created: time.Date(2023, 4, 16, 0, 0, 0, 0, time.Local),
							},
						}, nil
					},
				},
			},
			want: []model.Expense{
				{
					Id:      1,
					Amount:  33.5,
					Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
				},
				{
					Id:      1,
					Amount:  24.7,
					Created: time.Date(2023, 4, 16, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: false,
		},
		{
			name: "Got an empty array",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					FindAllFn: func() ([]model.Expense, error) {
						return []model.Expense{}, nil
					},
				},
			},
			want:    []model.Expense{},
			wantErr: false,
		},
		{
			name: "Got an error",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					FindAllFn: func() ([]model.Expense, error) {
						return nil, errors.New("error finding expenses")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				Repository: tt.fields.repository,
			}
			got, err := uc.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpenseUseCase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpenseUseCase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpenseUseCaseSave(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
	}
	type args struct {
		expense *model.Expense
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Expense
		wantErr bool
	}{
		{
			name: "given a expense, then save with success",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return false
					},
					SaveFn: func(e *model.Expense) (*model.Expense, error) {
						return e, nil
					},
				},
			},
			args: args{
				expense: &model.Expense{
					Id:      1,
					Amount:  100,
					Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
				},
			},
			want: &model.Expense{
				Id:      1,
				Amount:  100,
				Created: time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "given a expense, when try to save in database, then get error",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return false
					},
					SaveFn: func(e *model.Expense) (*model.Expense, error) {
						return nil, errors.ErrUnsupported
					},
				},
			},
			args: args{
				expense: &model.Expense{
					Id: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "given a expense, when the id is undefined, then get error",
			args: args{
				expense: &model.Expense{
					Id: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "given a expense, when exists in database, then get error",
			args: args{
				expense: &model.Expense{
					Id:     1,
					Amount: 100,
				},
			},
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return true
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				Repository: tt.fields.repository,
			}
			got, err := uc.Save(tt.args.expense)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpenseUseCase.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpenseUseCase.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpenseUseCase_Update(t *testing.T) {
	type fields struct {
		Repository port.ExpenseRepository
	}
	type args struct {
		expense *model.Expense
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Expense
		wantErr bool
	}{
		{
			name: "given a expense, update in database with success",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return true
					},
					UpdateFn: func(e *model.Expense) (*model.Expense, error) {
						return e, nil
					},
				},
			},
			args: args{
				expense: &model.Expense{
					Id:     1,
					Amount: 200,
				},
			},
			want: &model.Expense{
				Id:     1,
				Amount: 200,
			},
			wantErr: false,
		},
		{
			name: "given a expense, when the expense doesn't exists in database, then get error",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return false
					},
				},
			},
			args: args{
				expense: &model.Expense{
					Id:     1,
					Amount: 200,
				},
			},
			wantErr: true,
		},
		{
			name: "given a expense, when get an error on update in database, then get error",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return true
					},
					UpdateFn: func(e *model.Expense) (*model.Expense, error) {
						return nil, errors.ErrUnsupported
					},
				},
			},
			args: args{
				expense: &model.Expense{
					Id:     1,
					Amount: 200,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				Repository: tt.fields.Repository,
			}
			got, err := uc.Update(tt.args.expense)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpenseUseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpenseUseCase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpenseUseCase_Delete(t *testing.T) {
	type fields struct {
		Repository port.ExpenseRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "given an id, then delete item with success",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return true
					},
					DeleteFn: func(i int) error {
						return nil
					},
				},
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "given an id, when the item doesn't exist in database, then get error",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return false
					},
				},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
		{
			name: "given an id, when get an error on delete item, then get error",
			fields: fields{
				Repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(i int) bool {
						return true
					},
					DeleteFn: func(i int) error {
						return errors.ErrUnsupported
					},
				},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				Repository: tt.fields.Repository,
			}
			if err := uc.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ExpenseUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
