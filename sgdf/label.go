package sgdf

import (
	"github.com/tinque/totem/contact"
)

var labelCodeMapping = map[int][]contact.Label{
	110: []contact.Label{contact.LabelLouveteauJeannette},
	120: []contact.Label{contact.LabelScoutGuide},
	130: []contact.Label{contact.LabelPionnierCaravelle},
	140: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelCompagnon},
	170: []contact.Label{contact.LabelFarfadet},

	213: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelChefCheftaineLouveteauJeannette, contact.LabelChefCheftaine},
	223: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelChefCheftaineScoutGuide, contact.LabelChefCheftaine},
	233: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelChefCheftainePionnierCaravelle, contact.LabelChefCheftaine},
	240: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelAccompagnateurCompagnon},
	270: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelResponsableFarfadet},
	271: []contact.Label{contact.LabelParentFarfadet},
	300: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelBureau},
	302: []contact.Label{contact.LabelEquipeDeGroupe},
	307: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelBureau},
	309: []contact.Label{contact.LabelEquipeDeGroupe, contact.LabelBureau},
	330: []contact.Label{contact.LabelEquipeDeGroupe},
}

func getLabel(code int) []contact.Label {
	if labels, ok := labelCodeMapping[code]; ok {
		return labels
	}
	return nil
}
