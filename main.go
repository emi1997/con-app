package main

import (
	// "fmt"
	"github.com/urfave/cli"
	"os"
	"log"
)



var app = cli.NewApp()

func info () {
	app.Name = "console-app"
	app.Usage = "Storing school topics with elasticsearch"
	app.Author = "Julia Kettl"
	app.Version = "1.0.0"
}

func main() {
	info()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}