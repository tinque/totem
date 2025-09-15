package contact

import "slices"

// MergeContact merges the source contact into the destination contact.
// The source contact's data completes or replaces the destination's data
// according to the merge logic defined for each field.
func (c *Contact) MergeContact(source *Contact) {
	if source == nil {
		return
	}

	// CodeAdherant: keeps the destination value, merges only if empty
	if c.CodeAdherant == "" && source.CodeAdherant != "" {
		c.CodeAdherant = source.CodeAdherant
	}

	// FirstName: keeps the destination value, merges only if empty
	if c.FirstName == "" && source.FirstName != "" {
		c.FirstName = source.FirstName
	}

	// LastName: keeps the destination value, merges only if empty
	if c.LastName == "" && source.LastName != "" {
		c.LastName = source.LastName
	}

	// Emails: merges the maps, source emails complete the destination
	if source.Emails != nil {
		if c.Emails == nil {
			c.Emails = make(map[EmailType]string)
		}
		for emailType, email := range source.Emails {
			if email != "" && c.Emails[emailType] == "" {
				c.Emails[emailType] = email
			}
		}
	}

	// Birthday: keeps the destination value, merges only if nil
	if c.Birthday == nil && source.Birthday != nil {
		c.Birthday = source.Birthday
	}

	// Address: keeps the destination value, merges only if empty
	if c.Address == "" && source.Address != "" {
		c.Address = source.Address
	}

	// City: keeps the destination value, merges only if empty
	if c.City == "" && source.City != "" {
		c.City = source.City
	}

	// ZipCode: keeps the destination value, merges only if empty
	if c.ZipCode == "" && source.ZipCode != "" {
		c.ZipCode = source.ZipCode
	}

	// Country: keeps the destination value, merges only if empty
	if c.Country == "" && source.Country != "" {
		c.Country = source.Country
	}

	// Phones: merges the maps, source phones complete the destination
	if source.Phones != nil {
		if c.Phones == nil {
			c.Phones = make(map[PhoneType]string)
		}
		for phoneType, phone := range source.Phones {
			if phone != "" && c.Phones[phoneType] == "" {
				c.Phones[phoneType] = phone
			}
		}
	}

	// Position: keeps the destination value, merges only if empty
	if c.Position == "" && source.Position != "" {
		c.Position = source.Position
	}

	// Labels: merges the slices, adds source labels not already present
	if source.Labels != nil {
		for _, label := range source.Labels {
			if !slices.Contains(c.Labels, label) {
				c.Labels = append(c.Labels, label)
			}
		}
	}
}

// MergeContacts creates a new contact by merging two existing contacts.
// The first contact is used as the base, the second contact completes missing data.
// The original contacts are not modified.
func MergeContacts(destination, source *Contact) *Contact {
	if destination == nil && source == nil {
		return nil
	}
	if destination == nil {
		return copyContact(source)
	}
	if source == nil {
		return copyContact(destination)
	}

	merged := copyContact(destination)
	merged.MergeContact(source)
	return merged
}

// copyContact creates a deep copy of a contact
func copyContact(c *Contact) *Contact {
	if c == nil {
		return nil
	}

	copied := &Contact{
		CodeAdherant: c.CodeAdherant,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Address:      c.Address,
		City:         c.City,
		ZipCode:      c.ZipCode,
		Country:      c.Country,
		Position:     c.Position,
	}

	// Copy birthday
	if c.Birthday != nil {
		birthday := *c.Birthday
		copied.Birthday = &birthday
	}

	// Copy emails
	if c.Emails != nil {
		copied.Emails = make(map[EmailType]string, len(c.Emails))
		for k, v := range c.Emails {
			copied.Emails[k] = v
		}
	}

	// Copy phones
	if c.Phones != nil {
		copied.Phones = make(map[PhoneType]string, len(c.Phones))
		for k, v := range c.Phones {
			copied.Phones[k] = v
		}
	}

	// Copy labels
	if c.Labels != nil {
		copied.Labels = make([]Label, len(c.Labels))
		copy(copied.Labels, c.Labels)
	}

	return copied
}
