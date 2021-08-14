package psql

import (
	"cuan-tracker/internal/pkg/config"
	"testing"
)

func TestConnect(t *testing.T) {
	cfg := config.GetConfig("./config.yml")
	p := New(
		cfg.Database.Address,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Driver,
	)
	p.Connect()
	p.Ping()
}
