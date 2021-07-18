package app

import (
	"cuan-tracker/internal/pkg/config"
	"cuan-tracker/internal/pkg/logger"
	"cuan-tracker/pkg/psql"
)

const (
	ProcessedStatus      = "processed"
	FailedStatus         = "failed"
	JSONUnmarshalDetails = "unmarshal_json"
	DBError              = "db_error"
)

var db *psql.Psql

func Init() {
	logger.Info("app.Init", "app initialized", nil)
	initDB()
}

func Close() {
	db.Close()
}

func initDB() {
	cfg := config.GetConfig("config.yml")
	db = psql.New(
		cfg.Database.Address,
		cfg.Database.DBName,
		cfg.Database.User,
		cfg.Database.Pass,
	)
	db.Connect()
	db.Ping()
}
