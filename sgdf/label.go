package sgdf

import (
	"github.com/tinque/totem/contact"
)

var labelCodeMapping = map[int][]contact.Label{
	110: {contact.LabelLouveteauJeannette},
	120: {contact.LabelScoutGuide},
	130: {contact.LabelPionnierCaravelle},
	140: {contact.LabelEquipeDeGroupe, contact.LabelCompagnon},
	170: {contact.LabelFarfadet},

	213: {contact.LabelEquipeDeGroupe, contact.LabelChefCheftaineLouveteauJeannette, contact.LabelChefCheftaine},
	223: {contact.LabelEquipeDeGroupe, contact.LabelChefCheftaineScoutGuide, contact.LabelChefCheftaine},
	233: {contact.LabelEquipeDeGroupe, contact.LabelChefCheftainePionnierCaravelle, contact.LabelChefCheftaine},
	240: {contact.LabelEquipeDeGroupe, contact.LabelAccompagnateurCompagnon},
	270: {contact.LabelEquipeDeGroupe, contact.LabelResponsableFarfadet},
	271: {contact.LabelParentFarfadet},
	300: {contact.LabelEquipeDeGroupe, contact.LabelBureau},
	302: {contact.LabelEquipeDeGroupe},
	307: {contact.LabelEquipeDeGroupe, contact.LabelBureau},
	309: {contact.LabelEquipeDeGroupe, contact.LabelBureau},
	330: {contact.LabelEquipeDeGroupe},
}

func getLabel(code int) []contact.Label {
	if labels, ok := labelCodeMapping[code]; ok {
		return labels
	}
	return nil
}
