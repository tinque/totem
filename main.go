package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tinque/totem/address"
	"github.com/tinque/totem/contact"
	"github.com/tinque/totem/gmail"
	"github.com/tinque/totem/parser"
	"github.com/tinque/totem/sgdf"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	capitalizer := cases.Title(language.French, cases.Compact)

	for row := range rows {

		c := contact.Contact{}

		if v, ok := row["IndividuCivilite.CodeAdherent"]; ok {
			c.CodeAdherant = v
		}
		if v, ok := row["Individu.Prenom"]; ok {
			c.FirstName = capitalizer.String(v)
		}

		if v, ok := row["Individu.Nom"]; ok {
			c.LastName = capitalizer.String(v)
		}

		if v, ok := row["Individu.CourrielPersonnel"]; ok {
			c.SetEmail(contact.EmailPersonal, v)
		}

		if v, ok := row["Individu.CourrielDédiéSGDF"]; ok {
			c.SetEmail(contact.EmailDedicatedSGDF, v)
		}

		if v, ok := row["Individu.DateNaissance"]; ok {
			if v != "" {
				t, perr := time.Parse("02/01/2006", v)
				if perr != nil {
					fmt.Fprintf(os.Stderr, "error parsing date %q: %v\n", v, perr)
				} else {
					c.Birthday = &t
				}
			}
		}

		if v, ok := row["Individu.Adresse.Ligne1"]; ok {
			c.Address = address.FormatLine(v)
		}

		if v, ok := row["Individu.Adresse.Ligne2"]; ok {
			if c.Address != "" && v != "" {
				c.Address += "\n"
			}
			if v != c.Address && v != "" {
				c.Address += address.FormatLine(v)
			}
		}

		if v, ok := row["Individu.Adresse.Ligne3"]; ok {
			if c.Address != "" && v != "" {
				c.Address += "\n"
			}
			if v != c.Address && v != "" {
				c.Address += address.FormatLine(v)
			}
		}

		if v, ok := row["Individu.Adresse.CodePostal"]; ok {
			c.ZipCode = v
		}

		if v, ok := row["Individu.Adresse.Municipalite"]; ok {
			c.City = capitalizer.String(v)
		}

		if v, ok := row["Individu.Adresse.Pays"]; ok {
			c.Country = capitalizer.String(v)
		}

		if v, ok := row["Individu.TelephoneDomicile"]; ok {
			if v != "" {
				c.SetPhone(contact.PhoneHome, v)
			}
		}

		if v, ok := row["Individu.TelephonePortable1"]; ok {
			if v != "" {
				c.SetPhone(contact.PhoneMobile1, v)
			}
		}

		if v, ok := row["Individu.TelephonePortable2"]; ok {
			if v != "" {
				c.SetPhone(contact.PhoneMobile2, v)
			}
		}

		if v, ok := row["Individu.TelephoneBureau"]; ok {
			if v != "" {
				c.SetPhone(contact.PhoneWork, v)
			}
		}

		if v, ok := row["IndividuCivilite.NomCourt"]; ok {
			var gender sgdf.Gender
			switch v {
			case "M.":
				gender = sgdf.GenderMale
			case "P.":
				gender = sgdf.GenderMale
			case "Mme":
				gender = sgdf.GenderFemale
			}

			if v, ok := row["Fonction.Code"]; ok {
				i, err := strconv.Atoi(v)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error parsing position code %q: %v\n", v, err)
				} else {
					c.Position = string(sgdf.GetPosition(i, gender))
				}
			}
		}

		c.AddLabel(contact.LabelTest)

		cList = append(cList, c)
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
