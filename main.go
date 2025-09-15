package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tinque/totem/contact"
	"github.com/tinque/totem/gmail"
	"github.com/tinque/totem/parser"
	"github.com/tinque/totem/sgdf"
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

	cList := []contact.Contact{}
	for row := range rows {
		c, err := sgdf.ExtractIntranetContact(row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error extracting contact: %v\n", err)
			continue
		}

		cList = append(cList, c...)
	}

	w := csv.NewWriter(os.Stdout)

	csvContent := [][]string{}
	csvContent = append(csvContent, gmail.CSVHeader)
	for _, c := range cList {
		csvContent = append(csvContent, gmail.CSVContact(c))
	}

	outFile := path + ".csv"
	of, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("error creating output file %q: %v", outFile, err)
	}
	defer of.Close()

	// réutilise le writer mais vers le fichier
	w = csv.NewWriter(of)
	if err := w.WriteAll(csvContent); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	fmt.Fprintln(os.Stderr, "wrote", outFile)

}
