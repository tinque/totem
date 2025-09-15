package contact

import (
	"testing"
)

func TestAddLabel(t *testing.T) {
	c := &Contact{}
	c.AddLabel(LabelTest)
	if !c.HasLabel(LabelTest) {
		t.Errorf("LabelTest should be present after AddLabel")
	}
	// Test idempotence
	c.AddLabel(LabelTest)
	if len(c.Labels) != 1 {
		t.Errorf("LabelTest should not be duplicated")
	}
}

func TestHasLabel(t *testing.T) {
	c := &Contact{}
	if c.HasLabel(LabelTest) {
		t.Errorf("LabelTest should not be present initially")
	}
	c.AddLabel(LabelTest)
	if !c.HasLabel(LabelTest) {
		t.Errorf("LabelTest should be present after AddLabel")
	}
}

func TestRemoveLabel(t *testing.T) {
	c := &Contact{}
	c.AddLabel(LabelTest)
	c.RemoveLabel(LabelTest)
	if c.HasLabel(LabelTest) {
		t.Errorf("LabelTest should be removed")
	}
	// Remove non-existent label
	c.RemoveLabel(LabelTest)
	if len(c.Labels) != 0 {
		t.Errorf("Labels should remain empty after removing non-existent label")
	}
}
