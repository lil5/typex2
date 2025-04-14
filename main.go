package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lil5/typex2/internal/generate/dart"
	"github.com/lil5/typex2/internal/generate/kotlin"
	"github.com/lil5/typex2/internal/generate/rust"
	"github.com/lil5/typex2/internal/generate/swift"
	"github.com/lil5/typex2/internal/generate/typescript"
	"github.com/lil5/typex2/internal/read"
	"github.com/lil5/typex2/internal/utils"
	"github.com/lil5/typex2/internal/write"
	"github.com/lil5/typex2/tools"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "typex2",
		Usage:       "Convert go structs to other language types",
		Description: "Useful for generating types from golang json RestAPI projects for a frontend to ingest.\n\nExample:\ntypex2 -l kotlin -i . -o ./classes.kotlin",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "typescript",
				Aliases:     []string{"l", "language"},
				Usage:       "Language to generate to [typescript, dart, kotlin, swift, rust]",
				DefaultText: "typescript",
			},
			&cli.PathFlag{
				Name:      "input",
				Value:     ".",
				Aliases:   []string{"i"},
				Usage:     "Input directory to read go types from",
				TakesFile: false,
				Required:  true,
			},
			&cli.PathFlag{
				Name:      "output",
				Aliases:   []string{"o"},
				Usage:     "Output path & file to write translated types to",
				TakesFile: true,
				Required:  true,
			},
		},
		ArgsUsage: "path",
		Action: func(c *cli.Context) error {
			lang := c.String("lang")
			in := c.Path("input")
			out := c.Path("output")

			if len(lang) == 0 {
				fmt.Println(tools.NoEntry + " No language given")
				os.Exit(1)
			}

			mustDirExists(in)
			mustCheckOutFile(out, lang)

			pkg, err := read.LoadPackage(in)

			if err != nil {
				fmt.Printf(tools.NoEntry+" Loading packages failed: %v\n", err)
				return err
			}

			st := read.MapPackage(pkg)

			s, err := runLanguage(st, lang)

			if err != nil {
				fmt.Printf(tools.NoEntry + " Generate language failed")
				return err
			}

			err = write.FileWriter(out, s.String())
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

func mustDirExists(path string) {
	stat, err := os.Stat(path)

	if err != nil || !stat.IsDir() {
		fmt.Println(tools.NoEntry + " Path does not exist")
		os.Exit(1)
	}
}

func mustCheckOutFile(out string, lang string) {
	var suffix string
	switch lang {
	case utils.Typescript:
		suffix = ".ts"
	case utils.Dart:
		suffix = ".dart"
	case utils.Swift:
		suffix = ".swift"
	case utils.Kotlin:
		suffix = ".kotlin"
	case utils.Rust:
		suffix = ".rs"
	default:
		fmt.Println(tools.NoEntry + " Incorrect Language given")
		os.Exit(1)
	}
	if !strings.HasSuffix(out, suffix) {
		fmt.Println(tools.NoEntry + " Incorrect extension given to output file\nRequired extension: " + suffix)
		os.Exit(1)
	}

	mustDirExists(filepath.Dir(out))
}

func runLanguage(st *utils.StructMap, lang string) (*strings.Builder, error) {
	var err error
	var s *strings.Builder
	switch lang {
	case utils.Typescript:
		s, err = typescript.GenerateTypescript(st)
	case utils.Dart:
		s, err = dart.GenerateDart(st)
	case utils.Swift:
		s, err = swift.GenerateSwift(st)
	case utils.Kotlin:
		s, err = kotlin.GenerateKotlin(st)
	case utils.Rust:
		s, err = rust.GenerateRust(st)
	default:
		fmt.Println(tools.NoEntry + " Incorrect Language given")
		os.Exit(1)
	}

	return s, err
}
