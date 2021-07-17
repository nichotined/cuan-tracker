package psql

type Psql struct {
	addr     string
	database string
}

func New(addr string, database string) *Psql {
	return &Psql{addr: addr, database: database}
}

func (p *Psql) Query() {

}
