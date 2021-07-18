package psql

import "testing"

func TestConnect(t *testing.T) {
	p := New("127.0.0.1:5432", "cuan_tracker", "postgres", "")
	p.Connect()
	p.Ping()
}
