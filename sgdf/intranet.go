package sgdf

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tinque/totem/address"
	"github.com/tinque/totem/contact"
	"github.com/tinque/totem/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var capitalizer = cases.Title(language.French, cases.Compact)

func ExtractIntranetContact(row parser.Row) ([]contact.Contact, error) {
	var contacts []contact.Contact

	mainContact, err := extractIntranetMainContact(row)
	if err != nil {
		return nil, err
	}
	contacts = append(contacts, *mainContact)

	for i := 1; i <= 3; i++ {
		legalGuardianContact, err := extractIntranetLegalGardianContact(row, i)
		if err != nil {
			return nil, err
		}
		if legalGuardianContact != nil {
			contacts = append(contacts, *legalGuardianContact)
		}
	}

	return contacts, nil
}

func extractIntranetMainContact(row parser.Row) (*contact.Contact, error) {
	c := contact.Contact{}

	if v, ok := row["IndividuCivilite.CodeAdherent"]; ok {
		c.MemberCode = v
		c.AddLabel(contact.LabelAdherent)
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
		var gender Gender
		switch v {
		case "M.":
			gender = GenderMale
		case "P.":
			gender = GenderMale
		case "Mme":
			gender = GenderFemale
		}

		if v, ok := row["Fonction.Code"]; ok {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing position code %q: %v\n", v, err)
			} else {
				c.Position = string(getPosition(i, gender))
				labels := getLabel(i)
				for _, l := range labels {
					c.AddLabel(l)
				}
			}
		}
	}

	now := time.Now()
	c.UpdatedAt = &now

	return &c, nil
}

func extractIntranetLegalGardianContact(row parser.Row, index int) (*contact.Contact, error) {
	c := contact.Contact{}
	if v, ok := row[fmt.Sprintf("RepresentantLegal%dCivilite.NomCourt", index)]; ok {
		if v == "" {
			return nil, nil
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.CodeAdherent", index)]; ok {
		c.MemberCode = v
	}
	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Prenom", index)]; ok {
		c.FirstName = capitalizer.String(v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Nom", index)]; ok {
		c.LastName = capitalizer.String(v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.CourrielPersonnel", index)]; ok {
		c.SetEmail(contact.EmailPersonal, v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.CourrielDédiéSGDF", index)]; ok {
		c.SetEmail(contact.EmailDedicatedSGDF, v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.Ligne1", index)]; ok {
		c.Address = address.FormatLine(v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.Ligne2", index)]; ok {
		if c.Address != "" && v != "" {
			c.Address += "\n"
		}
		if v != c.Address && v != "" {
			c.Address += address.FormatLine(v)
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.Ligne3", index)]; ok {
		if c.Address != "" && v != "" {
			c.Address += "\n"
		}
		if v != c.Address && v != "" {
			c.Address += address.FormatLine(v)
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.CodePostal", index)]; ok {
		c.ZipCode = v
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.Municipalite", index)]; ok {
		c.City = capitalizer.String(v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.Adresse.Pays", index)]; ok {
		c.Country = capitalizer.String(v)
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.TelephoneDomicile", index)]; ok {
		if v != "" {
			c.SetPhone(contact.PhoneHome, v)
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.TelephonePortable1", index)]; ok {
		if v != "" {
			c.SetPhone(contact.PhoneMobile1, v)
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.TelephonePortable2", index)]; ok {
		if v != "" {
			c.SetPhone(contact.PhoneMobile2, v)
		}
	}

	if v, ok := row[fmt.Sprintf("RepresentantLegal%d.TelephoneBureau", index)]; ok {
		if v != "" {
			c.SetPhone(contact.PhoneWork, v)
		}
	}
	c.AddLabel(contact.LabelParent)
	now := time.Now()
	c.UpdatedAt = &now

	return &c, nil
}
