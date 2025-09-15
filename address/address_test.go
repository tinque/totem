package address

import "testing"

func TestFormatLine(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"31 AV DES KORRIGANS", "31 avenue des Korrigans"},
		{"10 RUE DE LA PAIX", "10 rue de la Paix"},
		{"5 BOULEVARD ST-MICHEL", "5 boulevard St-Michel"},
		{"12 av. du general de gaulle", "12 avenue du General de Gaulle"},
		{"7 bis rue de l'eglise", "7 bis rue de l'Eglise"},
		{"2 ter bd saint-michel", "2 ter boulevard Saint-Michel"},
		{"15 RTE DE LA PLAGE", "15 route de la Plage"},
		{"20 CHEMIN DES MIMOSAS", "20 chemin des Mimosas"},
		{"8 ALL DES TILLEULS", "8 allée des Tilleuls"},
		{"14 PL DE LA REPUBLIQUE", "14 place de la Republique"},
		{"3 SQUARE JEAN MOULIN", "3 square Jean Moulin"},
		{"1 IMPASSE DU LAC", "1 impasse du Lac"},
		{"4 PASSE DU COMMERCE", "4 passage du Commerce"},
		{"6 CITE INTERNATIONALE", "6 cité Internationale"},
		{"9 QUAI DE LA GARE", "9 quai de la Gare"},
		{"11 ROND-POINT DES CHAMPS", "11 rond-point des Champs"},
		{"13 VOIE ROMAINE", "13 voie Romaine"},
		{"17 PROMENADE DES ANGLAIS", "17 promenade des Anglais"},
		{"18 TERRASSE DU PORT", "18 terrasse du Port"},
		{"19 COUR DE L'ÉCOLE", "19 cour de l'École"},
		{"21 PLACE DU MARCHÉ", "21 place du Marché"},
		{"22 BOULEVARD DE LA LIBERTÉ", "22 boulevard de la Liberté"},
		{"", ""},
	}

	for _, c := range cases {
		got := FormatLine(c.in)
		if got != c.want {
			t.Fatalf("FormatLine(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}
