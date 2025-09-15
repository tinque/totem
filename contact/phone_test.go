package contact

import (
	"testing"
)

func TestSetAndGetPhone(t *testing.T) {
	c := &Contact{}
	c.SetPhone(PhoneMobile1, "0612345678")
	c.SetPhone(PhoneHome, "0145678901")

	if got := c.GetPhone(PhoneMobile1); got != "0612345678" {
		t.Errorf("GetPhone(PhoneMobile1) = %v, want %v", got, "0612345678")
	}
	if got := c.GetPhone(PhoneHome); got != "0145678901" {
		t.Errorf("GetPhone(PhoneHome) = %v, want %v", got, "0145678901")
	}
	if got := c.GetPhone(PhoneWork); got != "" {
		t.Errorf("GetPhone(PhoneWork) = %v, want empty string", got)
	}
}

func TestFirstPhone(t *testing.T) {
	c := &Contact{}
	if got := c.FirstPhone(); got != "" {
		t.Errorf("FirstPhone() = %v, want empty string", got)
	}

	c.SetPhone(PhoneHome, "0145678901")
	if got := c.FirstPhone(); got != "0145678901" {
		t.Errorf("FirstPhone() = %v, want %v", got, "0145678901")
	}

	c.SetPhone(PhoneMobile2, "0698765432")
	if got := c.FirstPhone(); got != "0698765432" {
		t.Errorf("FirstPhone() = %v, want %v", got, "0698765432")
	}

	c.SetPhone(PhoneMobile1, "0612345678")
	if got := c.FirstPhone(); got != "0612345678" {
		t.Errorf("FirstPhone() = %v, want %v", got, "0612345678")
	}
}
