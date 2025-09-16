package parser

import (
	"strings"
	"testing"
)

const sampleHTML = `<!doctype html>
<html>
  <body>
    <table>
      <tbody>
        <tr><td>Individu.Nom</td><td>Age</td></tr>
        <tr><td>Dupont</td><td>30</td></tr>
        <tr><td>Martin</td><td>25</td></tr>
      </tbody>
    </table>
  </body>
</html>`

func TestParseFromExcelHTMLReader(t *testing.T) {
	rows, err := FromExcelHTMLReader(strings.NewReader(sampleHTML))
	if err != nil {
		t.Fatalf("FromExcelHTMLReader failed: %v", err)
	}

	rowNumber := 0
	for row := range rows {

		if rowNumber == 0 && row["Individu.Nom"] != "Dupont" {
			t.Errorf("expected Dupont, got %q", row["Individu.Nom"])
		}
		if rowNumber == 0 && row["Age"] != "30" {
			t.Errorf("expected 30, got %q", row["Age"])
		}

		if rowNumber == 1 && row["Individu.Nom"] != "Martin" {
			t.Errorf("expected Martin, got %q", row["Individu.Nom"])
		}
		if rowNumber == 1 && row["Age"] != "25" {
			t.Errorf("expected 25, got %q", row["Age"])
		}

		rowNumber++
	}

	if rowNumber != 2 {
		t.Fatalf("expected 2 rows, got %d", rowNumber)
	}

}
