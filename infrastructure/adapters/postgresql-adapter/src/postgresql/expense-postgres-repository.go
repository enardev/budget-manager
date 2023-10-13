package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	customErrors "github.com/enaldo1709/budget-manager/domain/model/src/model/errors"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port"
	"github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter/src/postgresql/postgresconfig"
)

const (
	expensesTable = "expenses"
)

type ExpensePostgresAdapter struct {
	db     *sql.DB
	schema string
	table  string
}

func NewExpensePostgresAdapter(
	prop postgresconfig.PostgreSqlConnectionProperties, db *sql.DB) port.ExpenseRepository {
	return &ExpensePostgresAdapter{
		db:     db,
		schema: prop.Schema,
		table:  expensesTable,
	}
}

func (r *ExpensePostgresAdapter) Exists(id int) (bool, error) {
	query := fmt.Sprintf("select count(t.id) from %s.%s t where t.id = $1", r.schema, r.table)

	res, err := r.db.Query(query, id)
	if err != nil {
		log.Println("error: error executing query... ", err)
		return false, errors.Join(fmt.Errorf("error: error searching for expense... "), err)
	}
	var count int
	if res.Next() {
		if err = res.Scan(&count); err != nil {
			log.Println("error: error reading result... ", err)
			return false, errors.Join(fmt.Errorf("error: error reading exist result... "), err)
		}
	}

	return count > 0, nil
}

func (r *ExpensePostgresAdapter) FindByID(id int) (*model.Expense, error) {
	query := fmt.Sprintf("SELECT id, amount, created FROM %s.%s "+
		"WHERE id = $1", r.schema, r.table)

	res, err := r.db.Query(query, id)
	if err != nil {
		log.Println("error: error executing select query... ", err)
		return nil, errors.Join(fmt.Errorf("error: error searching for expense... "), err)
	}

	defer res.Close()

	if res.Next() {
		var retId int
		var amount float64
		var createdDate string
		err = res.Scan(&retId, &amount, &createdDate)
		if err != nil {
			log.Println("error: error building expense item... ", err)
			return nil, errors.Join(fmt.Errorf("error: error building expense item... "), err)
		}
		date, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			log.Println("error: error parsing created date... ", err)
			return nil, errors.Join(fmt.Errorf("error: error parsing created date... "), err)
		}
		return &model.Expense{Id: retId, Amount: amount, Created: date}, nil
	}

	return nil, customErrors.NewItemNotFoundError("expense")
}

func (r *ExpensePostgresAdapter) FindAll() ([]model.Expense, error) {
	query := fmt.Sprintf("SELECT id, amount, created FROM %s.%s", r.schema, r.table)
	res, err := r.db.Query(query)
	if err != nil {
		log.Println("error: error executing select query... ", err)
		return nil, errors.Join(fmt.Errorf("error: error searching for expenses... "), err)
	}

	expenses := []model.Expense{}

	defer res.Close()
	for res.Next() {
		var retId int
		var amount float64
		var createdDate string
		err = res.Scan(&retId, &amount, &createdDate)
		if err != nil {
			log.Println("error: error building expense item... ", err)
			return nil, errors.Join(fmt.Errorf("error: error building expense item... "), err)
		}
		date, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			log.Println("error: error parsing created date... ", err)
			return nil, errors.Join(fmt.Errorf("error: error parsing created date... "), err)
		}
		expenses = append(expenses, model.Expense{Id: retId, Amount: amount, Created: date})
	}

	return expenses, nil
}

func (r *ExpensePostgresAdapter) Save(e *model.Expense) (*model.Expense, error) {
	var nextVal int
	err := r.db.
		QueryRow(fmt.Sprintf("select nextval('%s.%s_id_seq'::regclass)", r.schema, r.table)).
		Scan(&nextVal)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("INSERT "+
		"INTO %s.%s (id, amount, created) "+
		"VALUES($1, $2, TO_TIMESTAMP($3, 'YYYY-MM-DD\"T\"HH24:MI:SS'))",
		r.schema, r.table)

	res, err := r.db.Exec(query, nextVal, e.Amount, e.Created.Format(time.RFC3339))
	if err != nil {
		log.Println("error: error executing insert query... ", err)
		return nil, errors.Join(fmt.Errorf("error: saving expense... "), err)
	}
	if nr, err := res.RowsAffected(); err != nil || nr == 0 {
		if err != nil {
			log.Println("error: error reading save result... ", err)
			return nil, errors.Join(fmt.Errorf("error: unknown save operation result... "), err)
		}
		log.Printf("error: error executing save query... %d items inserted\n", nr)
		return nil, fmt.Errorf("error: 0 items inserted on operation... ")
	}
	e.Id = nextVal
	return e, nil
}

func (r *ExpensePostgresAdapter) Update(e *model.Expense) (*model.Expense, error) {
	query := fmt.Sprintf("UPDATE %s.%s SET amount=$1, "+
		"created=TO_TIMESTAMP($2, 'YYYY-MM-DD\"T\"HH24:MI:SS') WHERE id=$3", r.schema, r.table)

	res, err := r.db.Exec(query, e.Amount, e.Created.Format(time.RFC3339), e.Id)
	if err != nil {
		log.Println("error: error executing update query... ", err)
		return nil, errors.Join(fmt.Errorf("error: updating expense... "), err)
	}
	if nr, err := res.RowsAffected(); err != nil || nr == 0 {
		if err != nil {
			log.Println("error: error reading update result... ", err)
			return nil, errors.Join(fmt.Errorf("error: unknown update operation result... "), err)
		}
		log.Printf("error: error executing update query... %d items updated\n", nr)
		return nil, fmt.Errorf("error: 0 items updated on operation... ")
	}
	return e, nil
}

func (r *ExpensePostgresAdapter) Delete(id int) error {
	query := fmt.Sprintf("DELETE fROM %s.%s WHERE id=$1", r.schema, r.table)

	res, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("error: error executing delete query... ", err)
		return errors.Join(fmt.Errorf("error: deleting expense... "), err)
	}
	if nr, err := res.RowsAffected(); err != nil || nr == 0 {
		if err != nil {
			log.Println("error: error reading delete result... ", err)
			return errors.Join(fmt.Errorf("error: unknown delete operation result... "), err)
		}
		log.Printf("error: error executing delete query... %d items deleted\n", nr)
		return fmt.Errorf("error: 0 items deleted on operation... ")
	}

	return nil
}
