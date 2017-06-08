package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/ryosan-470/factordb-go/lib"
)

func main() {
	app := cli.NewApp()
	app.Name = "factordb"
	app.Usage = "The CLI for factordb.com"
	app.Version = "1.0.0"

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
	var number = ""
	if len(c.Args()) > 0 {
		number = c.Args().First()
	}

	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Your input is not number")
		os.Exit(1)
	}

	f := lib.FactorDB{Number: n}
	if err := f.Connect(); err != nil {
		log.Fatal("Connection Error")
	}

	factors, _ := f.GetFactorList()

	var output string

	var isJson = c.GlobalBool("json")
	if isJson {
		id, _ := f.GetId()
		status, _ := f.GetStatus()
		var fs []string
		for _, f := range factors {
			fs = append(fs, fmt.Sprintf("%d", f))
		}

		facs := fmt.Sprintf("[%s]", strings.Join(fs, ", "))
		output = fmt.Sprintf("{\"id\": \"https://factordb.com/?id=%s\", \"status\": \"%s\", \"factors\": %v}", id, status, facs)
	} else {
		output = strings.Trim(fmt.Sprintf("%v", factors), "[]")
	}

	fmt.Printf("%s\n", output)
	return nil
}
