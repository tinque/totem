package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tinque/totem/parser"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <path-to-intranet-extract>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}
	path := flag.Arg(0)

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()

	rows, err := parser.FromReader(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing file: %v\n", err)
		os.Exit(2)
	}

	for row := range rows {
		if v, ok := row["Individu.Adresse.Ligne2"]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("(missing Individu.Adresse.Ligne2)")
		}
	}
}
