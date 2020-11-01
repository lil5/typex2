package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lil5/typex2/internal/generate/typescript"
	"github.com/lil5/typex2/internal/read"
	"github.com/lil5/typex2/internal/write"
	"github.com/lil5/typex2/tools"

	"github.com/urfave/cli/v2"
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

			pkg, err := read.LoadPackage(path)

			if err != nil {
				fmt.Printf(tools.NoEntry+" Loading packages failed: %v\n", err)
				return err
			}

			st := read.MapPackage(pkg)

			s, _ := typescript.GenerateTypescript(st)

			err = write.FileWriter(path, "index.ts", s)
			if err != nil {
				fmt.Printf(tools.NoEntry+" Write to file unsuccessful: %v\n", err)
				return err
			}

			fmt.Println(tools.CheckMark + " Written to file successfully")

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
