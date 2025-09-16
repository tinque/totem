package parser

import (
	"fmt"
	"io"
	"iter"
	"log"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Exported type for row data
// Format: header -> cell value
type Row map[string]string

// FromExcelHTMLReader parses an intranet export (Excel HTML) and returns rows.
func FromExcelHTMLReader(r io.Reader) (iter.Seq[Row], error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("parse html: %w", err)
	}
	return data(doc)
}

func getCellText(cell *html.Node) string {
	if cell == nil || cell.FirstChild == nil {
		return ""
	}
	if cell.DataAtom != atom.Td && cell.DataAtom != atom.Th {
		log.Fatal("getCellText: not a td or th")
	}
	return cell.FirstChild.Data
}

func findChildByAtom(n *html.Node, a atom.Atom) *html.Node {
	for c := range n.Descendants() {
		if c.DataAtom == a {
			return c
		}
	}
	return nil
}

func headers(tbody *html.Node) ([]string, error) {
	if tbody == nil {
		return nil, fmt.Errorf("tbody is nil")
	}
	tr := findChildByAtom(tbody, atom.Tr)
	if tr == nil {
		return nil, fmt.Errorf("no header tr found in tbody")
	}
	var headers []string
	for c := range tr.ChildNodes() {
		if c.DataAtom == atom.Td || c.DataAtom == atom.Th {
			headers = append(headers, getCellText(c))
		}
	}
	return headers, nil
}

func row(tr *html.Node, headers []string) Row {
	data := make(Row)
	j := 0
	for c := range tr.ChildNodes() {
		if c.DataAtom == atom.Td || c.DataAtom == atom.Th {
			text := getCellText(c)
			if j < len(headers) {
				data[headers[j]] = text
			} else {
				data[fmt.Sprintf("extra_%d", j-len(headers)+1)] = text
			}
			j++
		}
	}
	return data
}

func data(doc *html.Node) (iter.Seq[Row], error) {
	if doc == nil || doc.FirstChild == nil {
		return nil, fmt.Errorf("invalid HTML document")
	}
	htmlDoc := findChildByAtom(doc, atom.Html)
	if htmlDoc == nil {
		return nil, fmt.Errorf("no html node found")
	}
	body := findChildByAtom(htmlDoc, atom.Body)
	if body == nil {
		return nil, fmt.Errorf("no body node found")
	}
	table := findChildByAtom(body, atom.Table)
	if table == nil {
		return nil, fmt.Errorf("no table node found")
	}
	tbody := findChildByAtom(table, atom.Tbody)
	if tbody == nil {
		return nil, fmt.Errorf("no tbody node found")
	}

	headers, err := headers(tbody)
	if err != nil {
		return nil, err
	}

	return func(yield func(Row) bool) {
		seenHeader := false
		for tr := range tbody.ChildNodes() {
			if tr.DataAtom != atom.Tr {
				continue
			}
			if !seenHeader {
				seenHeader = true
				continue
			}
			if !yield(row(tr, headers)) {
				return
			}
		}
	}, nil
}
