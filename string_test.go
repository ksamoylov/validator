package validator

import (
	"golang.org/x/exp/maps"
	"testing"
)

func TestStringPass(t *testing.T) {
	fields := make(Fields, 0)
	secondFields := make(Fields, 0)

	fields["key"] = "key"
	fields["map"] = "map"
	fields["value"] = "value"

	maps.Copy(fields, secondFields)

	fRules := NewFieldRules(10, 2)

	rule := NewRule([]string{"key", "map"}, "string", *fRules)
	model := NewModel(fields, []Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(ToStringErrors(validationErrors))
		return
	}
}

func TestStringFailedType(t *testing.T) {
	fields := make(Fields, 0)

	fields["key"] = 123

	fRules := NewFieldRules(10, 2)

	rule := NewRule([]string{"key", "map"}, "string", *fRules)
	model := NewModel(fields, []Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(ToStringErrors(validationErrors))
		return
	}
}

func TestStringFailedLength(t *testing.T) {
	fields := make(Fields, 0)

	fields["key"] = "value"
	fields["value"] = "key"

	fRules := NewFieldRules(4, 4)

	rule := NewRule([]string{"key", "value"}, "string", *fRules)
	model := NewModel(fields, []Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(ToStringErrors(validationErrors))
		return
	}
}
