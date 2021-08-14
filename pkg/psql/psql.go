package psql

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Psql interface {
	Connect()
	Ping()
	Close()
	GetDBUrl() string
	GetClient() *pg.DB
}

type psql struct {
	addr     string
	database string
	port     string
	user     string
	password string
	driver   string
	db       *pg.DB
}

func New(addr string, port string, database string, user string, password string, driver string) Psql {
	return &psql{
		addr:     addr,
		port:     port,
		database: database,
		user:     user,
		password: password,
		driver:   driver,
	}
}

func (p *psql) GetDBUrl() string {
	return fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=require", p.driver, p.user, p.password, p.addr, p.port, p.database)
}

func (p *psql) Connect() {
	opt, err := pg.ParseURL(p.GetDBUrl())
	if err != nil {
		panic(err)
	}
	p.db = pg.Connect(opt)
}

func (p *psql) Ping() {
	ctx := context.Background()
	if err := p.db.Ping(ctx); err != nil {
		panic(err)
	}
}

func (p *psql) Close() {
	p.db.Close()
}

func (p *psql) GetClient() *pg.DB {
	return p.db
}
