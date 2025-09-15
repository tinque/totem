package contact

import (
	"reflect"
	"testing"
	"time"
)

func TestContact_MergeContact(t *testing.T) {
	// Date de test
	birthday1 := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	birthday2 := time.Date(1995, 6, 15, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		destination *Contact
		source      *Contact
		expected    *Contact
	}{
		{
			name: "Fusion avec source nil",
			destination: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			source: nil,
			expected: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		{
			name: "Fusion de champs vides avec champs remplis",
			destination: &Contact{
				FirstName: "John",
				LastName:  "",
				City:      "",
			},
			source: &Contact{
				FirstName: "Jane", // Ne devrait pas écraser
				LastName:  "Doe",
				City:      "Paris",
				ZipCode:   "75000",
			},
			expected: &Contact{
				FirstName: "John",  // Conservé
				LastName:  "Doe",   // Ajouté
				City:      "Paris", // Ajouté
				ZipCode:   "75000", // Ajouté
			},
		},
		{
			name: "Fusion des emails",
			destination: &Contact{
				FirstName: "John",
				Emails: map[EmailType]string{
					EmailPersonal: "john@personal.com",
				},
			},
			source: &Contact{
				FirstName: "John",
				Emails: map[EmailType]string{
					EmailPersonal:      "jane@personal.com", // Ne devrait pas écraser
					EmailDedicatedSGDF: "john@sgdf.org",     // Devrait être ajouté
				},
			},
			expected: &Contact{
				FirstName: "John",
				Emails: map[EmailType]string{
					EmailPersonal:      "john@personal.com", // Conservé
					EmailDedicatedSGDF: "john@sgdf.org",     // Ajouté
				},
			},
		},
		{
			name: "Fusion des téléphones",
			destination: &Contact{
				FirstName: "John",
				Phones: map[PhoneType]string{
					PhoneMobile1: "0123456789",
				},
			},
			source: &Contact{
				FirstName: "John",
				Phones: map[PhoneType]string{
					PhoneMobile1: "9876543210", // Ne devrait pas écraser
					PhoneHome:    "0147258369", // Devrait être ajouté
				},
			},
			expected: &Contact{
				FirstName: "John",
				Phones: map[PhoneType]string{
					PhoneMobile1: "0123456789", // Conservé
					PhoneHome:    "0147258369", // Ajouté
				},
			},
		},
		{
			name: "Fusion des labels",
			destination: &Contact{
				FirstName: "John",
				Labels:    []Label{LabelAdherent, LabelParent},
			},
			source: &Contact{
				FirstName: "John",
				Labels:    []Label{LabelParent, LabelChefCheftaine}, // LabelParent déjà présent
			},
			expected: &Contact{
				FirstName: "John",
				Labels:    []Label{LabelAdherent, LabelParent, LabelChefCheftaine},
			},
		},
		{
			name: "Fusion des dates d'anniversaire",
			destination: &Contact{
				FirstName: "John",
				Birthday:  nil,
			},
			source: &Contact{
				FirstName: "John",
				Birthday:  &birthday1,
			},
			expected: &Contact{
				FirstName: "John",
				Birthday:  &birthday1,
			},
		},
		{
			name: "Conservation de la date d'anniversaire existante",
			destination: &Contact{
				FirstName: "John",
				Birthday:  &birthday1,
			},
			source: &Contact{
				FirstName: "John",
				Birthday:  &birthday2, // Ne devrait pas écraser
			},
			expected: &Contact{
				FirstName: "John",
				Birthday:  &birthday1, // Conservé
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.destination.MergeContact(tt.source)
			if !reflect.DeepEqual(tt.destination, tt.expected) {
				t.Errorf("MergeContact() got = %+v, want %+v", tt.destination, tt.expected)
			}
		})
	}
}

func TestMergeContacts(t *testing.T) {
	birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		destination *Contact
		source      *Contact
		expected    *Contact
	}{
		{
			name:        "Deux contacts nil",
			destination: nil,
			source:      nil,
			expected:    nil,
		},
		{
			name:        "Destination nil",
			destination: nil,
			source: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		{
			name: "Source nil",
			destination: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
			source: nil,
			expected: &Contact{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		{
			name: "Fusion normale sans modification des originaux",
			destination: &Contact{
				FirstName: "John",
				LastName:  "",
				Birthday:  &birthday,
				Emails: map[EmailType]string{
					EmailPersonal: "john@test.com",
				},
			},
			source: &Contact{
				FirstName: "Jane", // Ne devrait pas écraser dans le résultat
				LastName:  "Doe",
				City:      "Paris",
				Emails: map[EmailType]string{
					EmailDedicatedSGDF: "john@sgdf.org",
				},
			},
			expected: &Contact{
				FirstName: "John",  // Conservé de destination
				LastName:  "Doe",   // Ajouté de source
				City:      "Paris", // Ajouté de source
				Birthday:  &birthday,
				Emails: map[EmailType]string{
					EmailPersonal:      "john@test.com",
					EmailDedicatedSGDF: "john@sgdf.org",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Sauvegarde les états originaux pour vérifier qu'ils ne sont pas modifiés
			var origDest, origSrc *Contact
			if tt.destination != nil {
				origDest = copyContact(tt.destination)
			}
			if tt.source != nil {
				origSrc = copyContact(tt.source)
			}

			result := MergeContacts(tt.destination, tt.source)

			// Vérifie le résultat
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeContacts() got = %+v, want %+v", result, tt.expected)
			}

			// Vérifie que les contacts originaux n'ont pas été modifiés
			if tt.destination != nil && !reflect.DeepEqual(tt.destination, origDest) {
				t.Errorf("MergeContacts() modified destination contact: got = %+v, want %+v", tt.destination, origDest)
			}
			if tt.source != nil && !reflect.DeepEqual(tt.source, origSrc) {
				t.Errorf("MergeContacts() modified source contact: got = %+v, want %+v", tt.source, origSrc)
			}
		})
	}
}

func TestCopyContact(t *testing.T) {
	birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		original *Contact
		expected *Contact
	}{
		{
			name:     "Contact nil",
			original: nil,
			expected: nil,
		},
		{
			name: "Contact complet",
			original: &Contact{
				CodeAdherant: "12345",
				FirstName:    "John",
				LastName:     "Doe",
				Birthday:     &birthday,
				Address:      "123 Rue de la Paix",
				City:         "Paris",
				ZipCode:      "75000",
				Country:      "France",
				Position:     "Chef",
				Emails: map[EmailType]string{
					EmailPersonal:      "john@personal.com",
					EmailDedicatedSGDF: "john@sgdf.org",
				},
				Phones: map[PhoneType]string{
					PhoneMobile1: "0123456789",
					PhoneHome:    "0147258369",
				},
				Labels: []Label{LabelAdherent, LabelChefCheftaine},
			},
			expected: &Contact{
				CodeAdherant: "12345",
				FirstName:    "John",
				LastName:     "Doe",
				Birthday:     &birthday,
				Address:      "123 Rue de la Paix",
				City:         "Paris",
				ZipCode:      "75000",
				Country:      "France",
				Position:     "Chef",
				Emails: map[EmailType]string{
					EmailPersonal:      "john@personal.com",
					EmailDedicatedSGDF: "john@sgdf.org",
				},
				Phones: map[PhoneType]string{
					PhoneMobile1: "0123456789",
					PhoneHome:    "0147258369",
				},
				Labels: []Label{LabelAdherent, LabelChefCheftaine},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copyContact(tt.original)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("copyContact() got = %+v, want %+v", result, tt.expected)
			}

			// Vérifie que c'est bien une copie profonde
			if tt.original != nil && result != nil {
				// Modifie l'original et vérifie que la copie n'est pas affectée
				tt.original.FirstName = "Modified"
				if result.FirstName == "Modified" {
					t.Error("copyContact() did not create a deep copy - FirstName was modified")
				}

				// Test avec les maps
				if tt.original.Emails != nil {
					tt.original.Emails[EmailPersonal] = "modified@test.com"
					if result.Emails[EmailPersonal] == "modified@test.com" {
						t.Error("copyContact() did not create a deep copy - Emails map was modified")
					}
				}

				// Test avec les slices
				if len(tt.original.Labels) > 0 {
					tt.original.Labels[0] = LabelBureau
					if len(result.Labels) > 0 && result.Labels[0] == LabelBureau {
						t.Error("copyContact() did not create a deep copy - Labels slice was modified")
					}
				}

				// Test avec la date
				if tt.original.Birthday != nil {
					*tt.original.Birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
					if result.Birthday != nil && result.Birthday.Year() == 2000 {
						t.Error("copyContact() did not create a deep copy - Birthday was modified")
					}
				}
			}
		})
	}
}
