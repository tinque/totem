package contact

type PhoneType string

const (
	PhoneHome    PhoneType = "home"
	PhoneMobile1 PhoneType = "mobile1"
	PhoneMobile2 PhoneType = "mobile2"
	PhoneWork    PhoneType = "work"
)

func (c *Contact) GetPhone(pt PhoneType) string {
	if c.Phones == nil {
		return ""
	}
	return c.Phones[pt]
}

func (c *Contact) SetPhone(pt PhoneType, number string) {
	if c.Phones == nil {
		c.Phones = make(map[PhoneType]string)
	}
	c.Phones[pt] = number
}

func (c *Contact) FirstPhone() string {
	order := []PhoneType{PhoneMobile1, PhoneMobile2, PhoneHome, PhoneWork}
	for _, pt := range order {
		if num := c.GetPhone(pt); num != "" {
			return num
		}
	}
	return ""
}
