package contact

type EmailType string

const (
	EmailPersonal      EmailType = "Personal"
	EmailDedicatedSGDF EmailType = "DedicatedSGDF"
)

func (c *Contact) GetEmail(et EmailType) string {
	if c.Emails == nil {
		return ""
	}
	return c.Emails[et]
}

func (c *Contact) SetEmail(et EmailType, email string) {
	if c.Emails == nil {
		c.Emails = make(map[EmailType]string)
	}
	c.Emails[et] = email
}

func (c *Contact) FirstEmail() string {
	order := []EmailType{EmailPersonal, EmailDedicatedSGDF}
	for _, et := range order {
		if email := c.GetEmail(et); email != "" {
			return email
		}
	}
	return ""
}
