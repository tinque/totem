package contact

import (
	"reflect"
	"testing"
	"time"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		{"Identical strings", "hello", "hello", 0},
		{"Empty strings", "", "", 0},
		{"One empty string", "hello", "", 5},
		{"Empty to non-empty", "", "world", 5},
		{"Single character diff", "hello", "hallo", 1},
		{"Complete different", "abc", "xyz", 3},
		{"Insert operation", "cat", "cart", 1},
		{"Delete operation", "cart", "cat", 1},
		{"Multiple operations", "kitten", "sitting", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levenshteinDistance(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("levenshteinDistance(%q, %q) = %d, want %d", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestNormalizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal string", "John", "john"},
		{"With spaces", " John ", "john"},
		{"Mixed case", "JoHn", "john"},
		{"Multiple spaces", "  John  Doe  ", "john  doe"},
		{"Empty string", "", ""},
		{"Only spaces", "   ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeString(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAreNamesSimilar(t *testing.T) {
	tests := []struct {
		name     string
		name1    string
		name2    string
		expected bool
	}{
		{"Identical names", "John", "John", true},
		{"Case difference", "John", "john", true},
		{"Spacing difference", " John ", "John", true},
		{"One typo", "John", "Jon", true},
		{"One typo", "John", "Jhn", true},
		{"Three typos", "John", "Jn", false}, // Too many differences for short name
		{"Different names", "John", "Peter", false},
		{"Empty names", "", "John", false},
		{"Both empty", "", "", false},
		{"Long similar names", "Alexander", "Alexandre", true},
		{"Long different names", "Alexander", "Sebastian", false},
		{"Short names similar", "Jo", "Ja", false}, // Too short, distance=1 but minLength < 3
		{"Accented characters", "José", "Jose", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := areNamesSimilar(tt.name1, tt.name2)
			if result != tt.expected {
				t.Errorf("areNamesSimilar(%q, %q) = %t, want %t", tt.name1, tt.name2, result, tt.expected)
			}
		})
	}
}

func TestAreDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		contact1 *Contact
		contact2 *Contact
		expected bool
	}{
		{
			name:     "Nil contacts",
			contact1: nil,
			contact2: nil,
			expected: false,
		},
		{
			name:     "One nil contact",
			contact1: &Contact{FirstName: "John"},
			contact2: nil,
			expected: false,
		},
		{
			name: "Same MemberCode",
			contact1: &Contact{
				MemberCode: "12345",
				FirstName:  "John",
				LastName:   "Doe",
			},
			contact2: &Contact{
				MemberCode: "12345",
				FirstName:  "Jane", // Different name but same code
				LastName:   "Smith",
			},
			expected: true,
		},
		{
			name: "Different MemberCode",
			contact1: &Contact{
				MemberCode: "12345",
				FirstName:  "John",
				LastName:   "Doe",
			},
			contact2: &Contact{
				MemberCode: "67890",
				FirstName:  "John",
				LastName:   "Doe",
			},
			expected: false, // Different codes, even with same names
		},
		{
			name: "No MemberCode, identical names",
			contact1: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			contact2: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: true,
		},
		{
			name: "No MemberCode, similar names with typos",
			contact1: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			contact2: &Contact{
				FirstName: "Jon", // One typo
				LastName:  "Do",  // One typo
			},
			expected: true,
		},
		{
			name: "No MemberCode, very different names",
			contact1: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			contact2: &Contact{
				FirstName: "Peter",
				LastName:  "Smith",
			},
			expected: false,
		},
		{
			name: "Missing first name",
			contact1: &Contact{
				FirstName: "",
				LastName:  "Doe",
			},
			contact2: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: false, // Missing first name
		},
		{
			name: "Missing last name",
			contact1: &Contact{
				FirstName: "John",
				LastName:  "",
			},
			contact2: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: false, // Missing last name
		},
		{
			name: "Empty MemberCode, same names",
			contact1: &Contact{
				MemberCode: "",
				FirstName:  "John",
				LastName:   "Doe",
			},
			contact2: &Contact{
				MemberCode: "",
				FirstName:  "John",
				LastName:   "Doe",
			},
			expected: true, // Fall back to name comparison
		},
		{
			name: "One has MemberCode, other doesn't",
			contact1: &Contact{
				MemberCode: "12345",
				FirstName:  "John",
				LastName:   "Doe",
			},
			contact2: &Contact{
				MemberCode: "",
				FirstName:  "John",
				LastName:   "Doe",
			},
			expected: true, // Fall back to name comparison since one MemberCode is empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := areDuplicates(tt.contact1, tt.contact2)
			if result != tt.expected {
				t.Errorf("areDuplicates() = %t, want %t", result, tt.expected)
			}
		})
	}
}

func TestDeduplicateAndMergeContacts(t *testing.T) {
	updatedAt1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedAt2 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		contacts []Contact
		expected []Contact
	}{
		{
			name:     "Empty list",
			contacts: []Contact{},
			expected: []Contact{},
		},
		{
			name: "Single contact",
			contacts: []Contact{
				{FirstName: "John", LastName: "Doe"},
			},
			expected: []Contact{
				{FirstName: "John", LastName: "Doe"},
			},
		},
		{
			name: "No duplicates",
			contacts: []Contact{
				{FirstName: "John", LastName: "Doe"},
				{FirstName: "Jane", LastName: "Smith"},
				{FirstName: "Bob", LastName: "Johnson"},
			},
			expected: []Contact{
				{FirstName: "John", LastName: "Doe"},
				{FirstName: "Jane", LastName: "Smith"},
				{FirstName: "Bob", LastName: "Johnson"},
			},
		},
		{
			name: "Duplicates by MemberCode",
			contacts: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "Doe",
					City:       "Paris",
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
				{
					MemberCode: "12345",
					FirstName:  "Johnny", // Different name but same code
					LastName:   "Doe",
					ZipCode:    "75000", // Additional info
				},
			},
			expected: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "John", // Original preserved
					LastName:   "Doe",
					City:       "Paris",
					ZipCode:    "75000", // Merged from duplicate
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
			},
		},
		{
			name: "Duplicates by similar names",
			contacts: []Contact{
				{
					FirstName: "John",
					LastName:  "Doe",
					City:      "Paris",
					Emails: map[EmailType]string{
						EmailPersonal: "john@test.com",
					},
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
				{
					FirstName: "Jon", // Similar to John (typo)
					LastName:  "Doe",
					ZipCode:   "75000",
					Emails: map[EmailType]string{
						EmailDedicatedSGDF: "john@sgdf.org",
					},
				},
			},
			expected: []Contact{
				{
					FirstName: "John", // Original preserved
					LastName:  "Doe",
					City:      "Paris",
					ZipCode:   "75000", // Merged from duplicate
					Emails: map[EmailType]string{
						EmailPersonal:      "john@test.com",
						EmailDedicatedSGDF: "john@sgdf.org",
					},
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
			},
		},
		{
			name: "Multiple duplicates of same contact",
			contacts: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "Doe",
					City:       "Paris",
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "Doe",
					ZipCode:    "75000",
				},
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "Doe",
					Country:    "France",
				},
			},
			expected: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "Doe",
					City:       "Paris",
					ZipCode:    "75000",
					Country:    "France",
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
				},
			},
		},
		{
			name: "Merge with UpdatedAt priority",
			contacts: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "John",
					LastName:   "OldName",
					UpdatedAt:  &updatedAt1, // Older
				},
				{
					MemberCode: "12345",
					FirstName:  "Johnny",
					LastName:   "NewName",
					UpdatedAt:  &updatedAt2, // Newer
				},
			},
			expected: []Contact{
				{
					MemberCode: "12345",
					FirstName:  "Johnny",  // From newer contact
					LastName:   "NewName", // From newer contact
					UpdatedAt:  &updatedAt2,
				},
			},
		},
		{
			name: "Complex scenario with multiple duplicate groups",
			contacts: []Contact{
				// Group 1: John Doe duplicates
				{
					FirstName: "John",
					LastName:  "Doe",
					City:      "Paris",
				},
				{
					FirstName: "Jon", // Typo
					LastName:  "Doe",
					ZipCode:   "75000",
				},
				// Group 2: Jane Smith (no duplicates)
				{
					FirstName: "Jane",
					LastName:  "Smith",
					City:      "Lyon",
				},
				// Group 3: Bob Johnson duplicates by code
				{
					MemberCode: "99999",
					FirstName:  "Bob",
					LastName:   "Johnson",
				},
				{
					MemberCode: "99999",
					FirstName:  "Robert", // Different name but same code
					LastName:   "Johnson",
					Position:   "Chef",
				},
			},
			expected: []Contact{
				{
					FirstName: "John", // Original name preserved
					LastName:  "Doe",
					City:      "Paris",
					ZipCode:   "75000", // Merged
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
					City:      "Lyon",
				},
				{
					MemberCode: "99999",
					FirstName:  "Bob", // Original name preserved
					LastName:   "Johnson",
					Position:   "Chef", // Merged
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeduplicateAndMergeContacts(tt.contacts)

			// Check that we have the expected number of contacts
			if len(result) != len(tt.expected) {
				t.Errorf("DeduplicateAndMergeContacts() returned %d contacts, want %d", len(result), len(tt.expected))
				return
			}

			// For each expected contact, find a matching result
			for i, expected := range tt.expected {
				if i >= len(result) {
					t.Errorf("Missing contact at index %d", i)
					continue
				}

				if !reflect.DeepEqual(result[i], expected) {
					t.Errorf("Contact at index %d:\ngot = %+v\nwant = %+v", i, result[i], expected)
				}
			}
		})
	}
}
