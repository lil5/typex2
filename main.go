package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lil5/typex2/internal/read"

	"github.com/lil5/typex2/tools"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:      "typex2",
		Usage:     "Convert go structs to typescript interfaces",
		ArgsUsage: "path",
		Action: func(c *cli.Context) error {
			path := c.Args().First()

			if len(path) == 0 {
				fmt.Println(tools.NoEntry + " No path given")
				os.Exit(1)
			}

			if !checkIfDirExists(path) {
				fmt.Println(tools.NoEntry + " Path does not exist")
			}

			pkgs, err := read.LoadPackages(path)

			if err != nil {
				fmt.Println(tools.NoEntry + " No structs found")
				return err
			}

			structTypes := read.MapPackages(pkgs)

			for _, st := range structTypes {
				if st == nil {
					continue
				}

				for n, t := range *st {
					fmt.Printf("Name: %s\n", n)
					fmt.Printf("Type: %v\n", t)
				}
			}

			// fmt.Printf("firstPkg: %v\n", firstPkg.Types)

			// fmt.Println(tools.Checkbox + " Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func checkIfDirExists(path string) bool {
	stat, err := os.Stat(path)

	return err == nil && stat.IsDir()
}
