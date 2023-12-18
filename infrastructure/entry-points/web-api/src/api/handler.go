package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/enaldo1709/budget-manager/domain/usecase/src/usecase"
	"github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api/src/api/dto"
	webErrors "github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api/src/api/errors"
	"github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api/src/api/mappers"
	"github.com/enaldo1709/budget-manager/infrastructure/helpers/validation/src/validation"
	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	Validator      *validation.Validator
	ExpenseUseCase usecase.ExpenseUseCase
}

func (h BudgetHandler) FindExpense(gc *gin.Context) {
	ids := gc.Request.URL.Query().Get("id")

	if ids == "" {
		items, err := h.ExpenseUseCase.FindAll()
		if err != nil {
			responseErr(gc, err)
			return
		}
		gc.JSON(http.StatusOK, mappers.MapExpenseListToDto(items))
		return
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		responseErr(gc, err)
		return
	}

	expense, err := h.ExpenseUseCase.FindByID(id)
	if err != nil {
		responseErr(gc, err)
		return
	}

	gc.JSON(http.StatusOK, mappers.MapExpenseToDto(expense))
}

func (h BudgetHandler) SaveExpense(gc *gin.Context) {
	var expenseDto dto.ExpenseDTO
	if err := json.NewDecoder(gc.Request.Body).Decode(&expenseDto); err != nil {
		responseErr(gc, err)
		return
	}

	expense, err := mappers.MapDtoToExpense(&expenseDto)
	if err != nil {
		responseErr(gc, err)
		return
	}

	res, err := h.ExpenseUseCase.Save(expense)
	if err != nil {
		responseErr(gc, err)
		return
	}

	gc.JSON(http.StatusCreated, mappers.MapExpenseToDto(res))
}

func (h BudgetHandler) UpdateExpense(gc *gin.Context) {
	var expenseDto dto.ExpenseDTO
	if err := h.Validator.DecodeAndValidate(gc.Request.Body, &expenseDto); err != nil {
		responseErr(gc, err)
		return
	}

	expense, err := mappers.MapDtoToExpense(&expenseDto)
	if err != nil {
		responseErr(gc, err)
		return
	}

	res, err := h.ExpenseUseCase.Update(expense)
	if err != nil {
		responseErr(gc, err)
		return
	}

	gc.JSON(http.StatusCreated, mappers.MapExpenseToDto(res))
}

func (h BudgetHandler) DeleteExpense(gc *gin.Context) {
	ids := gc.Request.URL.Query().Get("id")

	if ids == "" {
		responseErr(gc, errors.New("error: required parameter 'id' not found"))
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		responseErr(gc, err)
		return
	}

	if err = h.ExpenseUseCase.Delete(id); err != nil {
		responseErr(gc, err)
		return
	}

	gc.Status(http.StatusNoContent)
}

func responseErr(gc *gin.Context, err error) {
	webErr := webErrors.MapError(err)
	gc.JSON(webErr.Code, webErr)
}
