package postgres

import (
	"database/sql"
	"fmt"
	"github.com/cobbinma/example-go-api/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DBClient interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	DB() *sql.DB
}

type dbClient struct {
	db *sqlx.DB
}

func NewDBClient() DBClient {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBName,
		config.DBUser,
		config.DBPassword)

	driver := "postgres"

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logrus.Fatalln("Could not open database: ", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return &dbClient{db: db}
}

func (dbc *dbClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dbc.db.Exec(query, args...)
}

func (dbc *dbClient) Ping() error {
	return dbc.db.Ping()
}

func (dbc *dbClient) DB() *sql.DB {
	return dbc.db.DB
}