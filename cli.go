package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jonathanporta/go-sprintly-api"
	"os"
)

var username string
var token string

var resource string
var command string

func init() {
	username = "jonathan@getpantheon.com"
	token = "TOKENHERE"

	resource = flag.Arg(0)
	command = flag.Arg(1)
}

func main() {

	api := SprintlyAPI.Create(username, token)
	products, _, _ := api.Product().List()

	app := cli.NewApp()
	app.Name = "sprintly-cli"
	app.Usage = "Update Sprintly items from the command line!"

	app.Commands = []cli.Command{
		{
			Name:    "product",
			Aliases: []string{"p"},
			Usage:   "options for sprintly products",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all products",
					Action: func(c *cli.Context) {
						PrintProducts(products)
					},
				},
			},
		},
	}

	app.Run(os.Args)

}

func PrintProducts(products []SprintlyAPI.Product) {
	for i := range products {
		fmt.Println(products[i].Name)
	}
}
