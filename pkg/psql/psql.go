package psql

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Psql struct {
	addr     string
	database string
	user     string
	password string
	db       *pg.DB
}

func New(addr string, database string, user string, password string) *Psql {
	return &Psql{
		addr:     addr,
		database: database,
		user:     user,
		password: password,
	}
}

func (p *Psql) Connect() {
	p.db = pg.Connect(&pg.Options{
		Addr:     p.addr,
		User:     p.user,
		Password: p.password,
		Database: p.database,
	})
}

func (p *Psql) Ping() {
	ctx := context.Background()
	if err := p.db.Ping(ctx); err != nil {
		panic(err)
	}
}
