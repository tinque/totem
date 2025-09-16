package contact

import "slices"

type Label string

const LabelAdherent Label = "Adhérent"
const LabelBureau Label = "Bureau"
const LabelChefCheftaine Label = "Chef-Cheftaine"
const LabelParent Label = "Parent"
const LabelEquipeDeGroupe Label = "Equipe de groupe"

const LabelFarfadet Label = "Farfadet"
const LabelParentFarfadet Label = "Parent Farfadet"
const LabelResponsableFarfadet Label = "Responsable Farfadet"

const LabelLouveteauJeannette Label = "Louveteau-Jeannette"
const LabelParentLouveteauJeannette Label = "Parent Louveteau-Jeannette"
const LabelChefCheftaineLouveteauJeannette Label = "Chef-Cheftaine Louveteau-Jeannette"

const LabelPionnierCaravelle Label = "Pionnier-Caravelle"
const LabelParentPionnierCaravelle Label = "Parent Pionnier-Caravelle"
const LabelChefCheftainePionnierCaravelle Label = "Chef-Cheftaine Pionnier-Caravelle"

const LabelScoutGuide Label = "Scout-Guide"
const LabelParentScoutGuide Label = "Parent Scout-Guide"
const LabelChefCheftaineScoutGuide Label = "Chef-Cheftaine Scout-Guide"

const LabelCompagnon Label = "Compagnon"
const LabelParentCompagnon Label = "Parent Compagnon"
const LabelAccompagnateurCompagnon Label = "Accompagnateur Compagnon"

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

func (c *Contact) ClearManagedLabels() {
	managedLabels := []Label{
		LabelAdherent,
		LabelBureau,
		LabelChefCheftaine,
		LabelParent,
		LabelEquipeDeGroupe,
		LabelFarfadet,
		LabelParentFarfadet,
		LabelResponsableFarfadet,
		LabelLouveteauJeannette,
		LabelParentLouveteauJeannette,
		LabelChefCheftaineLouveteauJeannette,
		LabelPionnierCaravelle,
		LabelParentPionnierCaravelle,
		LabelChefCheftainePionnierCaravelle,
		LabelScoutGuide,
		LabelParentScoutGuide,
		LabelChefCheftaineScoutGuide,
		LabelCompagnon,
		LabelParentCompagnon,
		LabelAccompagnateurCompagnon,
	}

	c.Labels = slices.DeleteFunc(c.Labels, func(l Label) bool {
		return slices.Contains(managedLabels, l)
	})
}

func (c *Contact) LabelsAsStrings() []string {
	strs := make([]string, len(c.Labels))
	for i, label := range c.Labels {
		strs[i] = string(label)
	}
	return strs
}
