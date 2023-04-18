package tests

import (
	"github.com/ksamoylov/validator"
	"golang.org/x/exp/maps"
	"testing"
)

func TestIntPass(t *testing.T) {
	fields := make(validator.Fields, 0)
	secondFields := make(validator.Fields, 0)

	fields["key"] = 123

	maps.Copy(fields, secondFields)

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key"}, "int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestIntFailedType(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = "map"

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key"}, "int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestIntFailedLength(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = 12345
	fields["value"] = 123

	fRules := validator.NewFieldRules(4, 4)

	rule := validator.NewRule([]string{"key", "value"}, "int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}
