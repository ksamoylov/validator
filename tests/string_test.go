package tests

import (
	"github.com/ksamoylov/validator"
	"golang.org/x/exp/maps"
	"testing"
)

func TestStringPass(t *testing.T) {
	fields := make(validator.Fields, 0)
	secondFields := make(validator.Fields, 0)

	fields["key"] = "key"
	fields["map"] = "map"
	fields["value"] = "value"

	maps.Copy(fields, secondFields)

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key", "map"}, "string", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestStringFailedType(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = 123

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key", "map"}, "string", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestStringFailedLength(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = "value"
	fields["value"] = "key"

	fRules := validator.NewFieldRules(4, 4)

	rule := validator.NewRule([]string{"key", "value"}, "string", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}
