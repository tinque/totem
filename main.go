package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

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

	w := csv.NewWriter(os.Stdout)

	cIList := contactFromIntranet(path)
	cGList := contactFromGmail("contacts.csv")

	cList := append(cIList, cGList...)
	cList = contact.DeduplicateAndMergeContacts(cList)

	// Set the updated at timestamp
	now := time.Now()
	for i := range cList {
		cList[i].UpdatedAt = &now
	}

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

func contactFromIntranet(path string) []contact.Contact {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()

	rows, err := parser.FromExcelHTMLReader(f)
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

	cList = contact.DeduplicateAndMergeContacts(cList)

	return cList
}

func contactFromGmail(path string) []contact.Contact {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening contacts.csv: %v", err)
	}
	defer f.Close()

	rows, err := parser.FromCSVReader(f)
	if err != nil {
		log.Fatalf("error parsing contacts.csv: %v", err)
	}

	cList := []contact.Contact{}
	for row := range rows {
		c, err := gmail.ExtractGmailContact(row)
		if err != nil {
			log.Printf("error extracting contact: %v", err)
			continue
		}

		// clear labels
		c.ClearManagedLabels()
		c.RemoveLabel(contact.Label("Adhérant"))
		c.RemoveLabel(contact.Label("* myContacts"))

		cList = append(cList, c)
	}

	cList = contact.DeduplicateAndMergeContacts(cList)

	return cList
}
