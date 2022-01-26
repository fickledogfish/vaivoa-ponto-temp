package main

import (
	"log"
	"os"

	"github.com/fickledogfish/vaivoa-timesheets-cli/cli"
)

func main() {
	app, err := cli.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if err = app.Configure(); err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
