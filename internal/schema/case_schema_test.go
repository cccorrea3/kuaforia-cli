package schema

import "testing"

func TestValidateRequiredFields(t *testing.T) {
	errs := Validate(map[string]any{"title": "test"})
	if len(errs) == 0 {
		t.Error("expected error for missing description")
	}
}

func TestValidateValid(t *testing.T) {
	errs := Validate(map[string]any{
		"title":       "Test",
		"description": "A test case",
		"priority":    "high",
		"status":      "draft",
		"tags":        []any{"tag1", "tag2"},
	})
	if len(errs) > 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestValidateInvalidPriority(t *testing.T) {
	errs := Validate(map[string]any{
		"title":       "Test",
		"description": "A test case",
		"priority":    "urgent",
	})
	found := false
	for _, e := range errs {
		if e.Field == "priority" {
			found = true
		}
	}
	if !found {
		t.Error("expected validation error for invalid priority")
	}
}
