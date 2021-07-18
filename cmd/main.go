package main

import (
	"cuan-tracker/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	defer app.Close()
	app.Init()
	cli := newCLI()
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
