package contact

import "slices"

// MergeContact merges the source contact into the destination contact.
// The merge strategy depends on UpdatedAt timestamps:
// - A contact with UpdatedAt is always considered newer than one without UpdatedAt
// - If source has UpdatedAt but destination doesn't, source's non-empty data replaces destination's data
// - If destination has UpdatedAt but source doesn't, source's data only fills empty fields
// - If both have UpdatedAt, the one with the more recent timestamp determines the strategy
// - If neither has UpdatedAt, source's data only fills empty fields (conservative merge)
func (c *Contact) MergeContact(source *Contact) {
	if source == nil {
		return
	}

	// Determine merge strategy based on UpdatedAt timestamps
	// A contact with UpdatedAt is always newer than one without
	sourceIsNewer := false
	if source.UpdatedAt != nil && c.UpdatedAt == nil {
		// Source has UpdatedAt but destination doesn't -> source is newer
		sourceIsNewer = true
	} else if c.UpdatedAt != nil && source.UpdatedAt == nil {
		// Destination has UpdatedAt but source doesn't -> destination is newer
		sourceIsNewer = false
	} else if c.UpdatedAt != nil && source.UpdatedAt != nil {
		// Both have UpdatedAt -> compare timestamps
		sourceIsNewer = source.UpdatedAt.After(*c.UpdatedAt)
	}
	// If both are nil, sourceIsNewer remains false (conservative merge)

	// CodeAdherant: merge based on strategy
	if sourceIsNewer && source.CodeAdherant != "" {
		c.CodeAdherant = source.CodeAdherant
	} else if c.CodeAdherant == "" && source.CodeAdherant != "" {
		c.CodeAdherant = source.CodeAdherant
	}

	// FirstName: merge based on strategy
	if sourceIsNewer && source.FirstName != "" {
		c.FirstName = source.FirstName
	} else if c.FirstName == "" && source.FirstName != "" {
		c.FirstName = source.FirstName
	}

	// LastName: merge based on strategy
	if sourceIsNewer && source.LastName != "" {
		c.LastName = source.LastName
	} else if c.LastName == "" && source.LastName != "" {
		c.LastName = source.LastName
	}

	// Emails: merge maps based on strategy
	if source.Emails != nil {
		if c.Emails == nil {
			c.Emails = make(map[EmailType]string)
		}
		for emailType, email := range source.Emails {
			if email != "" {
				if sourceIsNewer {
					c.Emails[emailType] = email
				} else if c.Emails[emailType] == "" {
					c.Emails[emailType] = email
				}
			}
		}
	}

	// Birthday: merge based on strategy
	if sourceIsNewer && source.Birthday != nil {
		c.Birthday = source.Birthday
	} else if c.Birthday == nil && source.Birthday != nil {
		c.Birthday = source.Birthday
	}

	// Address: merge based on strategy
	if sourceIsNewer && source.Address != "" {
		c.Address = source.Address
	} else if c.Address == "" && source.Address != "" {
		c.Address = source.Address
	}

	// City: merge based on strategy
	if sourceIsNewer && source.City != "" {
		c.City = source.City
	} else if c.City == "" && source.City != "" {
		c.City = source.City
	}

	// ZipCode: merge based on strategy
	if sourceIsNewer && source.ZipCode != "" {
		c.ZipCode = source.ZipCode
	} else if c.ZipCode == "" && source.ZipCode != "" {
		c.ZipCode = source.ZipCode
	}

	// Country: merge based on strategy
	if sourceIsNewer && source.Country != "" {
		c.Country = source.Country
	} else if c.Country == "" && source.Country != "" {
		c.Country = source.Country
	}

	// Phones: merge maps based on strategy
	if source.Phones != nil {
		if c.Phones == nil {
			c.Phones = make(map[PhoneType]string)
		}
		for phoneType, phone := range source.Phones {
			if phone != "" {
				if sourceIsNewer {
					c.Phones[phoneType] = phone
				} else if c.Phones[phoneType] == "" {
					c.Phones[phoneType] = phone
				}
			}
		}
	}

	// Position: merge based on strategy
	if sourceIsNewer && source.Position != "" {
		c.Position = source.Position
	} else if c.Position == "" && source.Position != "" {
		c.Position = source.Position
	}

	// Labels: always merge (add source labels not already present)
	if source.Labels != nil {
		for _, label := range source.Labels {
			if !slices.Contains(c.Labels, label) {
				c.Labels = append(c.Labels, label)
			}
		}
	}

	// UpdatedAt: keep the most recent timestamp
	if source.UpdatedAt != nil {
		if c.UpdatedAt == nil || source.UpdatedAt.After(*c.UpdatedAt) {
			c.UpdatedAt = source.UpdatedAt
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

	// Copy UpdatedAt
	if c.UpdatedAt != nil {
		updatedAt := *c.UpdatedAt
		copied.UpdatedAt = &updatedAt
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
