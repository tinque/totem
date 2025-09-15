package sgdf

type Position string

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

const (
	PositionFarfadet                    Position = "Farfadet"
	PositionLouveteau                   Position = "Louveteau"
	PositionJeannette                   Position = "Jeannette"
	PositionScout                       Position = "Scout"
	PositionGuide                       Position = "Guide"
	PositionPionnier                    Position = "Pionnier"
	PositionCaravelle                   Position = "Caravelle"
	PositionCompagnon                   Position = "Compagnon"
	PositionChargeMissionGroupe         Position = "Charge de mission du groupe"
	PositionParentAnimateurFarfadet     Position = "Parent animateur Farfadet"
	PositionAccompagnateurCompagnon     Position = "Accompagnateur Compagnon"
	PositionChefLouveteauJeannette      Position = "Chef Louveteau Jeannette"
	PositionCheftaineLouveteauJeannette Position = "Cheftaine Louveteau Jeannette"
	PositionChefScoutGuide              Position = "Chef Scout Guide"
	PositionCheftaineScoutGuide         Position = "Cheftaine Scout Guide"
	PositionChefPionnierCaravelle       Position = "Chef Pionnier Caravelle"
	PositionCheftainePionnierCaravelle  Position = "Cheftaine Pionnier Caravelle"
	PositionResponsableFarfadet         Position = "Responsable Farfadet"
	PositionSecretaireGroupe            Position = "Secretaire de groupe"
	PositionResponsableGroupe           Position = "Responsable de groupe"
	PositionTresorierGroupe             Position = "Trésorier de groupe"
	PositionAumonierGroupe              Position = "Aumonier de groupe"
)

var positionCodeMapping = map[int]map[Gender]Position{
	110: {
		GenderMale:   PositionLouveteau,
		GenderFemale: PositionJeannette,
	},
	120: {
		GenderMale:   PositionScout,
		GenderFemale: PositionGuide,
	},
	130: {
		GenderMale:   PositionPionnier,
		GenderFemale: PositionCaravelle,
	},
	140: {
		GenderMale:   PositionCompagnon,
		GenderFemale: PositionCompagnon,
	},
	170: {
		GenderMale:   PositionFarfadet,
		GenderFemale: PositionFarfadet,
	},
	213: {
		GenderMale:   PositionChefLouveteauJeannette,
		GenderFemale: PositionCheftaineLouveteauJeannette,
	},
	223: {
		GenderMale:   PositionChefScoutGuide,
		GenderFemale: PositionCheftaineScoutGuide,
	},
	233: {
		GenderMale:   PositionChefPionnierCaravelle,
		GenderFemale: PositionCheftainePionnierCaravelle,
	},
	240: {
		GenderMale:   PositionAccompagnateurCompagnon,
		GenderFemale: PositionAccompagnateurCompagnon,
	},
	270: {
		GenderMale:   PositionResponsableFarfadet,
		GenderFemale: PositionResponsableFarfadet,
	},
	271: {
		GenderMale:   PositionParentAnimateurFarfadet,
		GenderFemale: PositionParentAnimateurFarfadet,
	},
	300: {
		GenderMale:   PositionResponsableGroupe,
		GenderFemale: PositionResponsableGroupe,
	},
	302: {
		GenderMale: PositionAumonierGroupe,
	},
	307: {
		GenderMale:   PositionSecretaireGroupe,
		GenderFemale: PositionSecretaireGroupe,
	},
	309: {
		GenderMale:   PositionSecretaireGroupe,
		GenderFemale: PositionSecretaireGroupe,
	},
	330: {
		GenderMale:   PositionChargeMissionGroupe,
		GenderFemale: PositionChargeMissionGroupe,
	},
}

func getPosition(code int, gender Gender) Position {
	if pos, ok := positionCodeMapping[code][gender]; ok {
		return pos
	}
	return ""
}
