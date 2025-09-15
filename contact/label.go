package contact

import "slices"

type Label string

const LabelTest Label = "TestTemporaire"

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

func (c *Contact) ClearLabels() {
	c.Labels = nil
}

func (c *Contact) LabelsAsStrings() []string {
	strs := make([]string, len(c.Labels))
	for i, label := range c.Labels {
		strs[i] = string(label)
	}
	return strs
}
