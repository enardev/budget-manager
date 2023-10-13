package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/enaldo1709/budget-manager/domain/model/src/model"
	"github.com/enaldo1709/budget-manager/domain/model/src/model/port"
	"github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter/src/postgresql/postgresconfig"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	expensesSchema = "test"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestNewExpensePostgresAdapter(t *testing.T) {
	db, _ := NewMock()
	defer db.Close()

	type args struct {
		prop postgresconfig.PostgreSqlConnectionProperties
		db   *sql.DB
	}
	tests := []struct {
		name string
		args args
		want port.ExpenseRepository
	}{
		{
			name: "given properties and database connection, then get a success repository instance",
			args: args{
				prop: postgresconfig.PostgreSqlConnectionProperties{Schema: "test"},
				db:   db,
			},
			want: &ExpensePostgresAdapter{
				db:     db,
				schema: "test",
				table:  expensesTable,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExpensePostgresAdapter(tt.args.prop, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExpensePostgresAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expensePostgresRepository_Exists(t *testing.T) {
	query := fmt.Sprintf("[select count(t.id) from %s.%s t where t.id = $1]",
		expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          bool
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given an id, when the expense exists in database, then return true",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			want: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

				return db, mock
			},
		},
		{
			name: "given an id, when the expense doesn't exists in database, then return false",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			want: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

				return db, mock
			},
		},
		{
			name: "given an id, when check if the expense exists in database, then return error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			want:    false,
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnError(errors.ErrUnsupported)

				return db, mock
			},
		},
		{
			name: "given an id, when check if the expense exists in database " +
				"and get an invalid result, then return error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			want:    false,
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("test"))

				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			got, err := r.Exists(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.Exists() error = %v, wantErr %v",
					err, tt.wantErr)
			}
			if (got != tt.want) && !tt.wantErr {
				t.Errorf("expensePostgresRepository.Exists() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func Test_expensePostgresRepository_FindByID(t *testing.T) {
	query := fmt.Sprintf("[SELECT id, amount, created FROM %s.%s WHERE id = $1]",
		expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          *model.Expense
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given an id, then get a success expense response",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			want: &model.Expense{
				Id:      1,
				Amount:  150,
				Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
			},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, 150, "2023-04-12T8:22:15Z"))
				return db, mock
			},
		},
		{
			name: "given an id, when an error occur in database, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnError(errors.ErrUnsupported)
				return db, mock
			},
		},
		{
			name: "given an id, when get invalid values, then get an error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, "test", "2023-04-12T8:22:15Z"))
				return db, mock
			},
		},
		{
			name: "given an id, when get an invalid created field, then get an error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, 150, "test"))
				return db, mock
			},
		},
		{
			name: "given an id, when expense is not found, then get an error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}))
				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			got, err := r.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expensePostgresRepository.FindByID() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func Test_expensePostgresRepository_FindAll(t *testing.T) {
	query := fmt.Sprintf("SELECT id, amount, created FROM %s.%s", expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	tests := []struct {
		name          string
		fields        fields
		want          []model.Expense
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given an get all expenses request, then get all expenses",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			want: []model.Expense{
				{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
				{
					Id:      2,
					Amount:  230,
					Created: time.Date(2023, 4, 12, 8, 26, 43, 0, time.UTC),
				},
				{
					Id:      3,
					Amount:  485,
					Created: time.Date(2023, 4, 12, 8, 33, 12, 0, time.UTC),
				},
			},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, 510, "2023-04-12T8:22:15Z").
						AddRow(2, 230, "2023-04-12T8:26:43Z").
						AddRow(3, 485, "2023-04-12T8:33:12Z"))
				return db, mock
			},
		},
		{
			name: "given an get all expenses request, when get a database error, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WillReturnError(errors.ErrUnsupported)
				return db, mock
			},
		},
		{
			name: "given an get all expenses request, when expenses database is empty, " +
				"get a empty response",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			want:    []model.Expense{},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}))
				return db, mock
			},
		},
		{
			name: "given an get all expenses request, when get an invalid value, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, "test", "2023-04-12T8:22:15Z"))
				return db, mock
			},
		},
		{
			name: "given an get all expenses request, when get an invalid created date, " +
				"then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "amount", "created"}).
						AddRow(1, 510, "test"))
				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			got, err := r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expensePostgresRepository.FindAll() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func Test_expensePostgresRepository_Save(t *testing.T) {
	querySeq := fmt.
		Sprintf("[select nextval('%s.%s_id_seq'::regclass)]", expensesSchema, expensesTable)
	query := fmt.Sprintf("[INSERT "+
		"INTO %s.%s (id, amount, created) "+
		"VALUES($1, $2, TO_TIMESTAMP($3, 'YYYY\\-MM\\-DD\"T\"HH24:MI:SS'))]",
		expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	type args struct {
		e *model.Expense
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          *model.Expense
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given an expense, when save with success in database, then get a expense",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			want: &model.Expense{
				Id:      1,
				Amount:  510,
				Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
			},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(querySeq).
					WillReturnRows(sqlmock.NewRows([]string{"nextval"}).AddRow(1))
				mock.ExpectExec(query).
					WithArgs(1, 510.0, "2023-04-12T08:22:15Z").
					WillReturnResult(sqlmock.NewResult(1, 1))

				return db, mock
			},
		},
		{
			name: "given an expense, when get an error checking next id, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(querySeq).
					WillReturnError(errors.ErrUnsupported)

				return db, mock
			},
		},
		{
			name: "given an expense, when there is an error executing in database, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(querySeq).
					WillReturnRows(sqlmock.NewRows([]string{"nextval"}).AddRow(1))
				mock.ExpectExec(query).
					WithArgs(1, 510.0, "2023-04-12T08:22:15Z").
					WillReturnError(errors.ErrUnsupported)

				return db, mock
			},
		},
		{
			name: "given an expense, when get error on reading operation result, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(querySeq).
					WillReturnRows(sqlmock.NewRows([]string{"nextval"}).AddRow(1))
				mock.ExpectExec(query).
					WithArgs(1, 510.0, "2023-04-12T08:22:15Z").
					WillReturnResult(sqlmock.NewErrorResult(errors.ErrUnsupported))

				return db, mock
			},
		},
		{
			name: "given an expense, when get 0 rows affected, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectQuery(querySeq).
					WillReturnRows(sqlmock.NewRows([]string{"nextval"}).AddRow(1))
				mock.ExpectExec(query).
					WithArgs(1, 510.0, "2023-04-12T08:22:15Z").
					WillReturnResult(sqlmock.NewResult(1, 0))

				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			got, err := r.Save(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expensePostgresRepository.Save() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func Test_expensePostgresRepository_Update(t *testing.T) {
	query := fmt.Sprintf("[UPDATE %s.%s SET amount=$1, "+
		"created=TO_TIMESTAMP($2, 'YYYY\\-MM\\-DD\"T\"HH24:MI:SS') WHERE id=$3]",
		expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	type args struct {
		e *model.Expense
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          *model.Expense
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given a expense to update, when update with success, then get an expense",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			want: &model.Expense{
				Id:      1,
				Amount:  510,
				Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
			},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(510.0, "2023-04-12T08:22:15Z", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				return db, mock
			},
		},
		{
			name: "given an expense to update, when there is an error in database, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(510.0, "2023-04-12T08:22:15Z", 1).
					WillReturnError(errors.ErrUnsupported)

				return db, mock
			},
		},
		{
			name: "given an expense to update, when there is an error result, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(510.0, "2023-04-12T08:22:15Z", 1).
					WillReturnResult(sqlmock.NewErrorResult(errors.ErrUnsupported))

				return db, mock
			},
		},
		{
			name: "given an expense to update, when get 0 rows affected, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				e: &model.Expense{
					Id:      1,
					Amount:  510,
					Created: time.Date(2023, 4, 12, 8, 22, 15, 0, time.UTC),
				},
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(510.0, "2023-04-12T08:22:15Z", 1).
					WillReturnResult(sqlmock.NewResult(1, 0))

				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			got, err := r.Update(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expensePostgresRepository.Update() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func Test_expensePostgresRepository_Delete(t *testing.T) {
	query := fmt.Sprintf("[DELETE fROM %s.%s WHERE id=$1]", expensesSchema, expensesTable)
	type fields struct {
		schema string
		table  string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantErr       bool
		configSqlMock func() (*sql.DB, sqlmock.Sqlmock)
	}{
		{
			name: "given an id, when tries to delete expense, " +
				"then delete with success and get nil error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: false,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				return db, mock
			},
		},
		{
			name: "given an id, when tries to delete expense an get database error, then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(1).
					WillReturnError(errors.ErrUnsupported)

				return db, mock
			},
		},
		{
			name: "given an id, when tries to delete expense and get error reading result," +
				" then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(1).
					WillReturnResult(sqlmock.NewErrorResult(errors.ErrUnsupported))

				return db, mock
			},
		},
		{
			name: "given an id, when tries to delete expense and fail with 0 rows affected" +
				" then get error",
			fields: fields{
				schema: expensesSchema,
				table:  expensesTable,
			},
			args: args{
				id: 1,
			},
			wantErr: true,
			configSqlMock: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock := NewMock()

				mock.ExpectExec(query).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 0))

				return db, mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configSqlMock == nil {
				t.Errorf("mock function is not configured")
				return
			}

			db, mock := tt.configSqlMock()
			defer db.Close()

			r := &ExpensePostgresAdapter{
				db:     db,
				schema: tt.fields.schema,
				table:  tt.fields.table,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("expensePostgresRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}
