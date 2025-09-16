package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"iter"
)

// FromCSVReader parses un export CSV et retourne les lignes.
func FromCSVReader(r io.Reader) (iter.Seq[Row], error) {
	reader := csv.NewReader(r)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV header: %v", err)
	}

	return func(yield func(Row) bool) {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
			row := make(Row)
			for i, value := range record {
				if i < len(headers) {
					row[headers[i]] = value
				} else {
					row[fmt.Sprintf("extra_%d", i-len(headers)+1)] = value
				}
			}
			if !yield(row) {
				return
			}
		}
	}, nil
}
