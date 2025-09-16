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
		{
			name: "Fusion intelligente - source plus récent écrase les données",
			destination: &Contact{
				FirstName: "John",
				LastName:  "OldName",
				City:      "OldCity",
				UpdatedAt: func() *time.Time { t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal: "old@email.com",
				},
			},
			source: &Contact{
				FirstName: "Jane",
				LastName:  "NewName",
				City:      "NewCity",
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com",
					EmailDedicatedSGDF: "jane@sgdf.org",
				},
			},
			expected: &Contact{
				FirstName: "Jane",    // Écrasé car source plus récent
				LastName:  "NewName", // Écrasé car source plus récent
				City:      "NewCity", // Écrasé car source plus récent
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com", // Écrasé car source plus récent
					EmailDedicatedSGDF: "jane@sgdf.org", // Ajouté
				},
			},
		},
		{
			name: "Fusion conservative - destination plus récente conserve ses données",
			destination: &Contact{
				FirstName: "John",
				LastName:  "NewName",
				City:      "NewCity",
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal: "new@email.com",
				},
			},
			source: &Contact{
				FirstName: "Jane",
				LastName:  "OldName",
				City:      "OldCity",
				ZipCode:   "12345", // Nouveau champ
				UpdatedAt: func() *time.Time { t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "old@email.com",
					EmailDedicatedSGDF: "jane@sgdf.org", // Nouveau email
				},
			},
			expected: &Contact{
				FirstName: "John",    // Conservé car destination plus récente
				LastName:  "NewName", // Conservé car destination plus récente
				City:      "NewCity", // Conservé car destination plus récente
				ZipCode:   "12345",   // Ajouté car vide dans destination
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com", // Conservé car destination plus récente
					EmailDedicatedSGDF: "jane@sgdf.org", // Ajouté car absent dans destination
				},
			},
		},
		{
			name: "Source avec UpdatedAt vs destination sans UpdatedAt - source écrase",
			destination: &Contact{
				FirstName: "John",
				LastName:  "OldName",
				City:      "OldCity",
				UpdatedAt: nil, // Pas de date (ancien contact)
				Emails: map[EmailType]string{
					EmailPersonal: "old@email.com",
				},
			},
			source: &Contact{
				FirstName: "Jane",
				LastName:  "NewName",
				City:      "NewCity",
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com",
					EmailDedicatedSGDF: "jane@sgdf.org",
				},
			},
			expected: &Contact{
				FirstName: "Jane",    // Écrasé car source a UpdatedAt
				LastName:  "NewName", // Écrasé car source a UpdatedAt
				City:      "NewCity", // Écrasé car source a UpdatedAt
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com", // Écrasé car source a UpdatedAt
					EmailDedicatedSGDF: "jane@sgdf.org", // Ajouté
				},
			},
		},
		{
			name: "Destination avec UpdatedAt vs source sans UpdatedAt - merge conservateur",
			destination: &Contact{
				FirstName: "John",
				LastName:  "NewName",
				City:      "NewCity",
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal: "new@email.com",
				},
			},
			source: &Contact{
				FirstName: "Jane",
				LastName:  "OldName",
				City:      "OldCity",
				ZipCode:   "12345", // Nouveau champ
				UpdatedAt: nil,     // Pas de date (ancien contact)
				Emails: map[EmailType]string{
					EmailPersonal:      "old@email.com",
					EmailDedicatedSGDF: "jane@sgdf.org", // Nouveau email
				},
			},
			expected: &Contact{
				FirstName: "John",    // Conservé car destination a UpdatedAt
				LastName:  "NewName", // Conservé car destination a UpdatedAt
				City:      "NewCity", // Conservé car destination a UpdatedAt
				ZipCode:   "12345",   // Ajouté car vide dans destination
				UpdatedAt: func() *time.Time { t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(),
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com", // Conservé car destination a UpdatedAt
					EmailDedicatedSGDF: "jane@sgdf.org", // Ajouté car absent dans destination
				},
			},
		},
		{
			name: "Aucun contact n'a UpdatedAt - merge conservateur",
			destination: &Contact{
				FirstName: "John",
				LastName:  "Existing",
				UpdatedAt: nil,
				Emails: map[EmailType]string{
					EmailPersonal: "existing@email.com",
				},
			},
			source: &Contact{
				FirstName: "Jane",  // Ne devrait pas écraser
				LastName:  "New",   // Ne devrait pas écraser
				ZipCode:   "12345", // Devrait être ajouté
				UpdatedAt: nil,
				Emails: map[EmailType]string{
					EmailPersonal:      "new@email.com", // Ne devrait pas écraser
					EmailDedicatedSGDF: "jane@sgdf.org", // Devrait être ajouté
				},
			},
			expected: &Contact{
				FirstName: "John",     // Conservé (merge conservateur)
				LastName:  "Existing", // Conservé (merge conservateur)
				ZipCode:   "12345",    // Ajouté car vide
				UpdatedAt: nil,
				Emails: map[EmailType]string{
					EmailPersonal:      "existing@email.com", // Conservé (merge conservateur)
					EmailDedicatedSGDF: "jane@sgdf.org",      // Ajouté car absent
				},
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
	updatedAt := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

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
				MemberCode: "12345",
				FirstName:  "John",
				LastName:   "Doe",
				Birthday:   &birthday,
				UpdatedAt:  &updatedAt,
				Address:    "123 Rue de la Paix",
				City:       "Paris",
				ZipCode:    "75000",
				Country:    "France",
				Position:   "Chef",
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
				MemberCode: "12345",
				FirstName:  "John",
				LastName:   "Doe",
				Birthday:   &birthday,
				UpdatedAt:  &updatedAt,
				Address:    "123 Rue de la Paix",
				City:       "Paris",
				ZipCode:    "75000",
				Country:    "France",
				Position:   "Chef",
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

				// Test avec UpdatedAt
				if tt.original.UpdatedAt != nil {
					*tt.original.UpdatedAt = time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)
					if result.UpdatedAt != nil && result.UpdatedAt.Year() == 2000 {
						t.Error("copyContact() did not create a deep copy - UpdatedAt was modified")
					}
				}
			}
		})
	}
}
