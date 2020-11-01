package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lil5/typex2/internal/generate/typescript"
	"github.com/lil5/typex2/internal/read"
	"github.com/lil5/typex2/internal/utils"
	"github.com/lil5/typex2/internal/write"
	"github.com/lil5/typex2/tools"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "typex2",
		Usage: "Convert go structs to typescript interfaces",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Value:   "typescript",
				Aliases: []string{"l", "language"},
				Usage:   "Language to generate to",
			},
		},
		ArgsUsage: "path",
		Action: func(c *cli.Context) error {
			path := c.Args().First()
			lang := c.String("lang")

			if len(path) == 0 {
				fmt.Println(tools.NoEntry + " No path given")
				os.Exit(1)
			}

			if len(lang) == 0 {
				fmt.Println(tools.NoEntry + " No language given")
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

			var fname string
			var s *string
			switch lang {
			case utils.Typescript:
				s, err = typescript.GenerateTypescript(st)
				fname = "index.ts"
			default:
				fmt.Println(tools.NoEntry + " Incorrect Language given")
				os.Exit(1)
			}

			if err != nil {
				fmt.Printf(tools.NoEntry + " Generate language failed")
				return err
			}

			err = write.FileWriter(path, fname, s)
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
