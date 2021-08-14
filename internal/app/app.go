package app

import (
	"cuan-tracker/internal/pkg/config"
	"cuan-tracker/pkg/psql"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	ProcessedStatus      = "processed"
	FailedStatus         = "failed"
	JSONUnmarshalDetails = "unmarshal_json"
	DBError              = "db_error"
)

var db psql.Psql

func Init() {
	initDB()
}

func Close() {
	db.Close()
}

func setupConnection() {
	cfg := config.GetConfig("./config.yml")
	db = psql.New(
		cfg.Database.Address,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Driver,
	)
}

func initDB() {
	setupConnection()
	db.Connect()
	db.Ping()
}

func MigrateDBUp() {
	initDB()
	dbase, err := sql.Open("postgres", db.GetDBUrl())
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(dbase, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func MigrateDBDown() {
	initDB()
	dbase, err := sql.Open("postgres", db.GetDBUrl())
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(dbase, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Down()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
