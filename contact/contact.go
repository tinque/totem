package contact

import (
	"time"
)

type Contact struct {
	CodeAdherant string
	FirstName    string
	LastName     string
	Emails       map[EmailType]string // Emails par type
	Birthday     *time.Time
	Address      string
	City         string
	ZipCode      string
	Country      string
	Phones       map[PhoneType]string // Numéros de téléphone par type
	Position     string
	Labels       []Label
	UpdatedAt    time.Time
}
