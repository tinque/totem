package contact

import (
	"strings"
)

// levenshteinDistance calculates the Levenshtein distance between two strings
func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// Create a matrix for dynamic programming
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		matrix[0][j] = j
	}

	// Fill the matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(s1)][len(s2)]
}

// normalizeString normalizes a string for comparison by removing extra spaces and converting to lowercase
func normalizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// areNamesSimilar checks if two names are similar enough to be considered duplicates
// Returns true if the names are identical or have a Levenshtein distance <= 2
func areNamesSimilar(name1, name2 string) bool {
	if name1 == "" || name2 == "" {
		return false
	}

	norm1 := normalizeString(name1)
	norm2 := normalizeString(name2)

	// If names are identical after normalization, they're similar
	if norm1 == norm2 {
		return true
	}

	// Calculate Levenshtein distance
	distance := levenshteinDistance(norm1, norm2)

	// Consider names similar if distance is <= 2 (allows for 1-2 typos)
	// But only if the names are not too short (to avoid false positives)
	minLength := min(len(norm1), len(norm2))
	if minLength >= 3 && distance <= 2 {
		return true
	}

	return false
}

// areDuplicates checks if two contacts are duplicates based on the specified criteria
func areDuplicates(contact1, contact2 *Contact) bool {
	if contact1 == nil || contact2 == nil {
		return false
	}

	// Priority 1: Same MemberCode (if both have one)
	if contact1.MemberCode != "" && contact2.MemberCode != "" {
		return contact1.MemberCode == contact2.MemberCode
	}

	// Priority 2: Similar first and last names
	if contact1.FirstName != "" && contact1.LastName != "" &&
		contact2.FirstName != "" && contact2.LastName != "" {
		return areNamesSimilar(contact1.FirstName, contact2.FirstName) &&
			areNamesSimilar(contact1.LastName, contact2.LastName)
	}

	return false
}

func DeduplicateAndMergeContacts(contacts []Contact) []Contact {
	if len(contacts) <= 1 {
		return contacts
	}

	var result []Contact
	processed := make([]bool, len(contacts))

	for i, contact := range contacts {
		if processed[i] {
			continue
		}

		// Current contact becomes the base for merging
		mergedContact := copyContact(&contact)
		processed[i] = true

		// Look for duplicates of this contact
		for j := i + 1; j < len(contacts); j++ {
			if processed[j] {
				continue
			}

			if areDuplicates(&contact, &contacts[j]) {
				// Merge the duplicate into our base contact
				mergedContact.MergeContact(&contacts[j])
				processed[j] = true
			}
		}

		result = append(result, *mergedContact)
	}

	return result
}
