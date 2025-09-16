package gmail

import (
	"fmt"
	"strings"
	"time"

	"github.com/tinque/totem/contact"
	"github.com/tinque/totem/parser"
)

func extractCSVCustomField(label, value string, c *contact.Contact) {
	if label == "Code Adhérent" && value != "" {
		c.MemberCode = value
	}

	if label == "Dernière mise à jour" && value != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", value); err == nil {
			c.UpdatedAt = &t
		}
	}
}

func extractCSVEmail(label, value string, c *contact.Contact) {
	if c.Emails == nil {
		c.Emails = make(map[contact.EmailType]string)
	}
	if label == "Personnel" && value != "" {
		c.Emails[contact.EmailPersonal] = value
	}

	if label == "Dédié SGDF" && value != "" {
		c.Emails[contact.EmailDedicatedSGDF] = value
	}
}

func extractCSVPhone(label, value string, c *contact.Contact) {
	if c.Phones == nil {
		c.Phones = make(map[contact.PhoneType]string)
	}
	if label == "Mobile 1" && value != "" {
		c.Phones[contact.PhoneMobile1] = value
	}

	if label == "Mobile 2" && value != "" {
		c.Phones[contact.PhoneMobile2] = value
	}

	if label == "Domicile" && value != "" {
		c.Phones[contact.PhoneHome] = value
	}

	if label == "Travail" && value != "" {
		c.Phones[contact.PhoneWork] = value
	}
}

func ExtractGmailContact(row parser.Row) (contact.Contact, error) {
	c := contact.Contact{}

	// Custom fields
	for i := 1; i <= 2; i++ {
		if v, ok := row[fmt.Sprintf("Custom Field %d - Value", i)]; ok {
			if l, ok := row[fmt.Sprintf("Custom Field %d - Label", i)]; ok {
				extractCSVCustomField(l, v, &c)
			}
		}
	}

	// Base information
	if v, ok := row["First Name"]; ok {
		c.FirstName = v
	}
	if v, ok := row["Last Name"]; ok {
		c.LastName = v
	}
	if v, ok := row["Organization Title"]; ok {
		c.Position = v
	}
	if v, ok := row["Birthday"]; ok {
		if v != "" {
			birthday, err := time.Parse("2006-01-02", v)
			if err == nil {
				c.Birthday = &birthday
			}
		}
	}
	if v, ok := row["Labels"]; ok {
		parts := strings.Split(v, " ::: ")
		labels := make([]contact.Label, 0, len(parts))
		for _, p := range parts {
			if p == "" {
				continue
			}
			labels = append(labels, contact.Label(p))
		}
		c.Labels = labels
	}

	// Emails and Phones
	for i := 1; i <= 2; i++ {
		if v, ok := row[fmt.Sprintf("E-mail %d - Value", i)]; ok {
			if l, ok := row[fmt.Sprintf("E-mail %d - Label", i)]; ok {
				extractCSVEmail(l, v, &c)
			}
		}
	}

	for i := 1; i <= 4; i++ {
		if v, ok := row[fmt.Sprintf("Phone %d - Value", i)]; ok {
			if l, ok := row[fmt.Sprintf("Phone %d - Label", i)]; ok {
				extractCSVPhone(l, v, &c)
			}
		}
	}

	// Address information
	if v, ok := row["Address 1 - Street"]; ok {
		c.Address = v
	}
	if v, ok := row["Address 1 - Postal Code"]; ok {
		c.ZipCode = v
	}
	if v, ok := row["Address 1 - City"]; ok {
		c.City = v
	}
	if v, ok := row["Address 1 - Country"]; ok {
		c.Country = v
	}

	return c, nil

}
