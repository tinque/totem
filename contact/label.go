package contact

import "slices"

type Label string

const LabelTest Label = "test"

func (c *Contact) AddLabel(label Label) {
	if slices.Contains(c.Labels, label) {
		return // Déjà présent
	}
	c.Labels = append(c.Labels, label)
}

func (c *Contact) HasLabel(label Label) bool {
	return slices.Contains(c.Labels, label)
}

func (c *Contact) RemoveLabel(label Label) {
	c.Labels = slices.DeleteFunc(c.Labels, func(l Label) bool { return l == label })
}
