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
	intranetPath := flag.String("intranet", "", "Path to intranet extract file (required)")
	gmailPath := flag.String("gmail", "", "Path to Gmail contacts CSV file (optional)")
	outputPath := flag.String("out", "output.csv", "Path to output CSV file (optional)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -intranet <extract-intranet> [-gmail <contacts.csv>] [-out <output.csv>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *intranetPath == "" {
		fmt.Fprintln(os.Stderr, "The -intranet parameter is required.")
		flag.Usage()
		os.Exit(2)
	}

	cIList := contactFromIntranet(*intranetPath)
	cList := cIList
	if *gmailPath != "" {
		cGList := contactFromGmail(*gmailPath)
		cList = append(cList, cGList...)
	}
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

	of, err := os.Create(*outputPath)
	if err != nil {
		log.Fatalf("error creating output file %q: %v", *outputPath, err)
	}
	defer of.Close()

	w := csv.NewWriter(of)
	if err := w.WriteAll(csvContent); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	fmt.Fprintln(os.Stderr, "wrote", *outputPath)

}

func contactFromIntranet(path string) []contact.Contact {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()

	rows, err := parser.FromExcelHTMLReader(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(2)
	}

	cList := []contact.Contact{}
	for row := range rows {
		c, err := sgdf.ExtractIntranetContact(row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting contact: %v\n", err)
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
		log.Fatalf("Error opening contacts.csv: %v", err)
	}
	defer f.Close()

	rows, err := parser.FromCSVReader(f)
	if err != nil {
		log.Fatalf("Error parsing contacts.csv: %v", err)
	}

	cList := []contact.Contact{}
	for row := range rows {
		c, err := gmail.ExtractGmailContact(row)
		if err != nil {
			log.Printf("Error extracting contact: %v", err)
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
