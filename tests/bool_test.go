package tests

import (
	"github.com/ksamoylov/validator"
	"testing"
)

func TestBoolPass(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = true
	fields["value"] = false

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key", "value"}, "bool", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestBoolFailedType(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = "map"

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key"}, "bool", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}
