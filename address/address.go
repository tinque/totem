// Package address provides utilities for formatting French street addresses.
// It handles common abbreviations, applies proper French capitalization rules,
// and manages special cases like compound words and French articles.
package address

import (
	"strings"
	"unicode"
)

// streetAbbreviations maps common street type abbreviations to their full forms
var streetAbbreviations = map[string]string{
	"AV":         "avenue",
	"AV.":        "avenue",
	"AVENUE":     "avenue",
	"RUE":        "rue",
	"BD":         "boulevard",
	"BD.":        "boulevard",
	"BOULEVARD":  "boulevard",
	"RTE":        "route",
	"RTE.":       "route",
	"CHEMIN":     "chemin",
	"ALL":        "allée",
	"ALL.":       "allée",
	"PL":         "place",
	"PL.":        "place",
	"IMPASSE":    "impasse",
	"SQ":         "square",
	"SQUARE":     "square",
	"PASSE":      "passage",
	"PASSE.":     "passage",
	"CITE":       "cité",
	"QUAI":       "quai",
	"ROND-POINT": "rond-point",
	"PROMENADE":  "promenade",
	"TERRASSE":   "terrasse",
	"COUR":       "cour",
	"VOIE":       "voie",
}

// frenchSmallWords contains French articles and prepositions that should remain lowercase
// when not at the beginning of the address
var frenchSmallWords = map[string]bool{
	"de": true, "du": true, "des": true, "la": true, "le": true, "les": true,
	"l'": true, "d'": true, "sur": true, "au": true, "aux": true, "et": true,
}

// streetTypes contains street type words that should remain lowercase when not first
var streetTypes = map[string]bool{
	"avenue": true, "rue": true, "boulevard": true, "route": true, "chemin": true,
	"allée": true, "allee": true, "place": true, "impasse": true, "square": true,
	"passage": true, "cité": true, "cite": true, "quai": true, "rond-point": true,
	"promenade": true, "terrasse": true, "cour": true, "voie": true,
}

// FormatLine formats a French address line by normalizing abbreviations,
// applying proper capitalization rules, and handling special cases
func FormatLine(line string) string {
	if line == "" {
		return ""
	}

	// Step 1: Expand abbreviations
	expanded := expandAbbreviations(line)

	// Step 2: Apply proper capitalization
	return applyFrenchCapitalization(expanded)

}

// expandAbbreviations replaces common street abbreviations with their full forms
func expandAbbreviations(line string) string {
	tokens := strings.Fields(line)

	for i, token := range tokens {
		upperToken := removeTrailingPunctuation(strings.ToUpper(token))
		if fullForm, exists := streetAbbreviations[upperToken]; exists {
			tokens[i] = fullForm
		}
	}

	return strings.Join(tokens, " ")
}

// applyFrenchCapitalization applies proper French capitalization rules
func applyFrenchCapitalization(line string) string {
	line = strings.ToLower(line)
	tokens := strings.Fields(line)

	for i, token := range tokens {
		if isNumber(token) || token == "bis" || token == "ter" {
			continue // Keep numbers and ordinals as-is
		}

		tokens[i] = capitalizeToken(token, i == 0)
	}

	return strings.Join(tokens, " ")
}

// capitalizeToken capitalizes a token according to French address rules
func capitalizeToken(token string, isFirstToken bool) string {
	lowToken := strings.ToLower(token)

	// Keep small words and street types lowercase unless first token
	if !isFirstToken && (frenchSmallWords[lowToken] || streetTypes[lowToken]) {
		return lowToken
	}

	return capitalizeWithSpecialChars(token)
}

// capitalizeWithSpecialChars handles capitalization of words with hyphens and apostrophes
// Special rules:
// - Words separated by hyphens are each capitalized: "jean-claude" -> "Jean-Claude"
// - French articles after apostrophes remain lowercase: "l'avenue" -> "l'Avenue", "d'artagnan" -> "d'Artagnan"
// - Other words after apostrophes are capitalized: "qu'est-ce" -> "Qu'Est-Ce"
func capitalizeWithSpecialChars(word string) string {
	// Handle simple cases first
	if !strings.ContainsAny(word, "-'") {
		return capitalizeFirstLetter(word)
	}

	var result strings.Builder
	var currentWord strings.Builder

	for _, char := range word {
		if char == '-' || char == '\'' {
			// Process accumulated word before separator
			if currentWord.Len() > 0 {
				wordStr := currentWord.String()
				// Special case: French articles 'l' and 'd' remain lowercase after apostrophes
				if strings.ToLower(wordStr) == "l" || strings.ToLower(wordStr) == "d" {
					result.WriteString(strings.ToLower(wordStr))
				} else {
					result.WriteString(capitalizeFirstLetter(wordStr))
				}
				currentWord.Reset()
			}
			result.WriteRune(char)
		} else {
			currentWord.WriteRune(char)
		}
	}

	// Process final word part
	if currentWord.Len() > 0 {
		wordStr := currentWord.String()
		// Special case: French articles 'l' and 'd' remain lowercase after apostrophes
		if strings.ToLower(wordStr) == "l" || strings.ToLower(wordStr) == "d" {
			result.WriteString(strings.ToLower(wordStr))
		} else {
			result.WriteString(capitalizeFirstLetter(wordStr))
		}
	}

	return result.String()
}

// Helper functions

// removeTrailingPunctuation removes trailing punctuation from a token
func removeTrailingPunctuation(token string) string {
	return strings.TrimRight(token, ".,")
}

// isNumber checks if a token is numeric
func isNumber(token string) bool {
	if token == "" {
		return false
	}

	for i, char := range token {
		if i == 0 && (char == '+' || char == '-') {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// capitalizeFirstLetter capitalizes the first letter of a word
func capitalizeFirstLetter(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
