package contact

import (
	"testing"
)

func TestGetEmail(t *testing.T) {
	c := &Contact{}
	if got := c.GetEmail(PersonalEmail); got != "" {
		t.Errorf("expected empty string, got %q", got)
	}

	c.SetEmail(PersonalEmail, "john@example.com")
	if got := c.GetEmail(PersonalEmail); got != "john@example.com" {
		t.Errorf("expected 'john@example.com', got %q", got)
	}
}

func TestSetEmail(t *testing.T) {
	c := &Contact{}
	c.SetEmail(DedicatedSGDFEmail, "sgdf@example.com")
	if got := c.Emails[DedicatedSGDFEmail]; got != "sgdf@example.com" {
		t.Errorf("expected 'sgdf@example.com', got %q", got)
	}
}

func TestFirstEmail(t *testing.T) {
	c := &Contact{}
	if got := c.FirstEmail(); got != "" {
		t.Errorf("expected empty string, got %q", got)
	}

	c.SetEmail(DedicatedSGDFEmail, "sgdf@example.com")
	if got := c.FirstEmail(); got != "sgdf@example.com" {
		t.Errorf("expected 'sgdf@example.com', got %q", got)
	}

	c.SetEmail(PersonalEmail, "john@example.com")
	if got := c.FirstEmail(); got != "john@example.com" {
		t.Errorf("expected 'john@example.com', got %q", got)
	}
}
