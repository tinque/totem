package contact

import (
	"testing"
)

func TestAddLabel(t *testing.T) {
	c := &Contact{}
	c.AddLabel(LabelAdherent)
	if !c.HasLabel(LabelAdherent) {
		t.Errorf("LabelAdherent should be present after AddLabel")
	}
	// Test idempotence
	c.AddLabel(LabelAdherent)
	if len(c.Labels) != 1 {
		t.Errorf("LabelAdherent should not be duplicated")
	}
}

func TestHasLabel(t *testing.T) {
	c := &Contact{}
	if c.HasLabel(LabelAdherent) {
		t.Errorf("LabelAdherent should not be present initially")
	}
	c.AddLabel(LabelAdherent)
	if !c.HasLabel(LabelAdherent) {
		t.Errorf("LabelAdherent should be present after AddLabel")
	}
}

func TestRemoveLabel(t *testing.T) {
	c := &Contact{}
	c.AddLabel(LabelAdherent)
	c.RemoveLabel(LabelAdherent)
	if c.HasLabel(LabelAdherent) {
		t.Errorf("LabelAdherent should be removed")
	}
	// Remove non-existent label
	c.RemoveLabel(LabelAdherent)
	if len(c.Labels) != 0 {
		t.Errorf("Labels should remain empty after removing non-existent label")
	}
}
