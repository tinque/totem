package gmail

import (
	"fmt"
	"log"
	"strings"

	"github.com/tinque/totem/contact"
)

type csvField struct {
	Label string
	Value string
}

var CSVHeader = []string{
	"Name Prefix",
	"First Name",
	"Middle Name",
	"Last Name",
	"Name Suffix",
	"Phonetic First Name",
	"Phonetic Middle Name",
	"Phonetic Last Name",
	"Nickname",
	"File as",
	"E-mail 1 - Label",
	"E-mail 1 - Value",
	"E-mail 2 - Label",
	"E-mail 2 - Value",
	"Phone 1 - Label",
	"Phone 1 - Value",
	"Phone 2 - Label",
	"Phone 2 - Value",
	"Phone 3 - Label",
	"Phone 3 - Value",
	"Phone 4 - Label",
	"Phone 4 - Value",
	"Address 1 - Label",
	"Address 1 - Country",
	"Address 1 - Street",
	"Address 1 - Extended Address",
	"Address 1 - City",
	"Address 1 - Region",
	"Address 1 - Postal Code",
	"Address 1 - PO Box",
	"Organization Name",
	"Organization Title",
	"Organization Department",
	"Birthday",
	"Notes",
	"Labels",
	"Custom Field 1 - Value",
	"Custom Field 1 - Label",
	"Custom Field 2 - Value",
	"Custom Field 2 - Label",
	// "Relation 1 - Label",
	// "Relation 1 - Value",
	// "Relation 2 - Label",
	// "Relation 2 - Value",
	// "Relation 3 - Label",
	// "Relation 3 - Value",
}

func getHeaderIndex(header string) int {
	for i, h := range CSVHeader {
		if h == header {
			return i
		}
	}
	log.Fatalln("Header not found:", header)
	panic("unreachable")
}

func mapEmailsToCSV(row []string, c contact.Contact) {
	var fields []csvField

	if email, ok := c.Emails[contact.EmailPersonal]; ok {
		fields = append(fields, csvField{Label: "Personnel", Value: email})
	}

	if email, ok := c.Emails[contact.EmailDedicatedSGDF]; ok {
		fields = append(fields, csvField{Label: "Dédié SGDF", Value: email})
	}
	for i, field := range fields {
		if i >= 2 { // Limite à 2 emails maximum
			break
		}

		row[getHeaderIndex(fmt.Sprintf("E-mail %d - Label", i+1))] = field.Label
		row[getHeaderIndex(fmt.Sprintf("E-mail %d - Value", i+1))] = field.Value
	}
}

func mapPhonesToCSV(row []string, c contact.Contact) {
	var fields []csvField

	// Ordre de priorité : Mobile1, Mobile2, Home, Work
	phoneTypes := []struct {
		phoneType contact.PhoneType
		label     string
	}{
		{contact.PhoneMobile1, "Mobile 1"},
		{contact.PhoneMobile2, "Mobile 2"},
		{contact.PhoneHome, "Domicile"},
		{contact.PhoneWork, "Travail"},
	}

	for _, pt := range phoneTypes {
		if phone, ok := c.Phones[pt.phoneType]; ok {
			fields = append(fields, csvField{Label: pt.label, Value: phone})
		}
	}

	for i, field := range fields {
		if i >= 4 { // Limite à 4 téléphones maximum
			break
		}

		row[getHeaderIndex(fmt.Sprintf("Phone %d - Label", i+1))] = field.Label
		row[getHeaderIndex(fmt.Sprintf("Phone %d - Value", i+1))] = field.Value
	}
}

func CSVContact(c contact.Contact) []string {
	row := make([]string, len(CSVHeader))

	// Custom fields
	if c.MemberCode != "" {
		row[getHeaderIndex("Custom Field 1 - Label")] = "Code Adhérent"
		row[getHeaderIndex("Custom Field 1 - Value")] = c.MemberCode
	}

	if c.UpdatedAt != nil {
		row[getHeaderIndex("Custom Field 2 - Label")] = "Dernière mise à jour"
		row[getHeaderIndex("Custom Field 2 - Value")] = c.UpdatedAt.Format("2006-01-02 15:04:05")
	}

	// Base information
	row[getHeaderIndex("First Name")] = c.FirstName
	row[getHeaderIndex("Last Name")] = c.LastName
	row[getHeaderIndex("Organization Title")] = c.Position
	if c.Birthday != nil {
		row[getHeaderIndex("Birthday")] = c.Birthday.Format("2006-01-02")
	}
	row[getHeaderIndex("Labels")] = strings.Join(c.LabelsAsStrings(), " ::: ")

	// Emails and Phones
	mapEmailsToCSV(row, c)
	mapPhonesToCSV(row, c)

	// Address information
	row[getHeaderIndex("Address 1 - Label")] = "Domicile" // Address 1 - Label
	row[getHeaderIndex("Address 1 - Street")] = c.Address
	row[getHeaderIndex("Address 1 - Postal Code")] = c.ZipCode
	row[getHeaderIndex("Address 1 - City")] = c.City
	row[getHeaderIndex("Address 1 - Country")] = c.Country

	return row
}
