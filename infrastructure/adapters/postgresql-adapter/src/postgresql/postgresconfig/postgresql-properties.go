package postgresconfig

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSqlConnectionProperties struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	Schema   string `yaml:"schema"`
}

func CreateSqlConnection(properties PostgreSqlConnectionProperties) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		properties.Host, properties.Port, properties.User, properties.Password, properties.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(errors.Join(fmt.Errorf("error creating database connection -> %v", err), err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal("cannot connect with database... ", err)
	}

	return db
}
