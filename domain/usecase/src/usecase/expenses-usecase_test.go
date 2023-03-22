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

func TestExpenseUseCase_FindByID(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
	}
	type args struct {
		id string
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
					ExistsFn: func(s string) bool {
						return true
					},
					FindByIDFn: func(id string) (*model.Expense, error) {
						return &model.Expense{
							Id:     "31667aaf-8c90-4887-bf13-5c4598689656",
							Amount: 25.3,
							Date:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
						}, nil
					},
				},
			},
			args: args{id: "31667aaf-8c90-4887-bf13-5c4598689656"},
			want: &model.Expense{
				Id:     "31667aaf-8c90-4887-bf13-5c4598689656",
				Amount: 25.3,
				Date:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "given an id when the expense not exists then get an error",
			fields: fields{
				repository: &mocks.ExpenseRepositoryMock{
					ExistsFn: func(s string) bool {
						return false
					},
				},
			},
			args:    args{id: "31667aaf-8c90-4887-bf13-5c4598689656"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				repository: tt.fields.repository,
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

func TestExpenseUseCase_FindAll(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
		idGen      port.IdGenerator
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
								Id:     "04c41b3c-d877-430b-a451-8a662ea5b684",
								Amount: 33.5,
								Date:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
							},
							{
								Id:     "33dc1676-48f6-4a49-9d3c-bbb0959c4551",
								Amount: 24.7,
								Date:   time.Date(2023, 4, 16, 0, 0, 0, 0, time.Local),
							},
						}, nil
					},
				},
			},
			want: []model.Expense{
				{
					Id:     "04c41b3c-d877-430b-a451-8a662ea5b684",
					Amount: 33.5,
					Date:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local),
				},
				{
					Id:     "33dc1676-48f6-4a49-9d3c-bbb0959c4551",
					Amount: 24.7,
					Date:   time.Date(2023, 4, 16, 0, 0, 0, 0, time.Local),
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
				repository: tt.fields.repository,
				idGen:      tt.fields.idGen,
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

func TestExpenseUseCase_Save(t *testing.T) {
	type fields struct {
		repository port.ExpenseRepository
		idGen      port.IdGenerator
	}
	type args struct {
		expense model.Expense
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Expense
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := ExpenseUseCase{
				repository: tt.fields.repository,
				idGen:      tt.fields.idGen,
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
