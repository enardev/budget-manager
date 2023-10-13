package main

import (
	"log"

	configv2 "github.com/gookit/config/v2"

	"github.com/enaldo1709/budget-manager/domain/usecase/src/usecase"
	"github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter/src/postgresql"
	"github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter/src/postgresql/postgresconfig"
	"github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api/src/api"
	"github.com/enaldo1709/budget-manager/infrastructure/helpers/configutil/src/configutil"
)

func main() {
	configutil.LoadConfig()

	// PostgreSQL database configuration
	properties := postgresconfig.PostgreSqlConnectionProperties{}
	configv2.MapStruct("db.properties", &properties)

	//repository := postgresql.NewExpensePostgresAdapter(properties, db)
	db := postgresconfig.CreateSqlConnection(properties)
	repository := postgresql.NewExpensePostgresAdapter(properties, db)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("error: error closing database connection...", err)
		}
	}()

	// UseCase configuration
	expenseUseCase := usecase.ExpenseUseCase{Repository: repository}

	// Api configuration
	budgetHandler := api.BudgetHandler{
		ExpenseUseCase: expenseUseCase,
	}
	router := api.ConfigRouter(budgetHandler)
	router.Run()
}
