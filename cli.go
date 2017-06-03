package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "factordb"
	app.Usage = "The CLI for factordb.com"
	app.Version = "0.0.1"

	// Global option
	app.Flags = []cli.Flag{
		// --json
		cli.BoolFlag{
			Name:  "json",
			Usage: "Return response formated JSON",
		},
	}

	app.Action = callAction
	app.Run(os.Args)
}

func callAction(c *cli.Context) error {
	var isJson = c.GlobalBool("json")
	if isJson {
		fmt.Println("Response json type")
	}

	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	}

	fmt.Printf("Hi, I am receiving the number %s\n", paramFirst)
	return nil
}
