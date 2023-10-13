package mappers

import (
	"fmt"
	"log"
	"time"

	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	"github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api/src/api/dto"
)

func MapDtoToExpense(d *dto.ExpenseDTO) (*model.Expense, error) {
	created, err := time.Parse(time.RFC3339, d.Created)
	if err != nil {
		log.Println("error: error decoding created date... ", err)
		return nil, fmt.Errorf("invalid field 'created', must be ISO date format")
	}

	return &model.Expense{
		Id:      d.Id,
		Amount:  d.Amount,
		Created: created,
	}, nil
}

func MapExpenseToDto(e *model.Expense) *dto.ExpenseDTO {
	return &dto.ExpenseDTO{
		Id:      e.Id,
		Amount:  e.Amount,
		Created: e.Created.Format(time.RFC3339),
	}
}

func MapExpenseListToDto(d []model.Expense) []dto.ExpenseDTO {
	dtos := []dto.ExpenseDTO{}

	for _, e := range d {
		dtos = append(dtos, *MapExpenseToDto(&e))
	}

	return dtos
}
