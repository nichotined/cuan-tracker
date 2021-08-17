package main

import (
	"cuan-tracker/internal/app"
	"fmt"
)

func main() {
	fmt.Println("cuan-tracker server")
	defer app.Close()
	app.Init()
	cli := newCLI()
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
