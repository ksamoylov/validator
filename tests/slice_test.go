package tests

import (
	"github.com/ksamoylov/validator"
	"testing"
)

func TestSliceIntPass(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["slice"] = []int{1, 2, 3}

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"slice"}, "[]int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestSliceStringPass(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["slice"] = []string{"one", "two", "three"}

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"slice"}, "[]string", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Errorf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestSliceFailedType(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = 123
	fields["map"] = "123"

	fRules := validator.NewFieldRules(10, 2)

	rule := validator.NewRule([]string{"key", "map"}, "[]int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}

func TestSliceFailedLength(t *testing.T) {
	fields := make(validator.Fields, 0)

	fields["key"] = []int{1, 2, 3, 4, 5}
	fields["value"] = []int{1, 2, 3}

	fRules := validator.NewFieldRules(4, 4)

	rule := validator.NewRule([]string{"key", "value"}, "[]int", *fRules)
	model := validator.NewModel(fields, []validator.Rule{*rule})

	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		t.Skipf(validator.ToStringErrors(validationErrors))
		return
	}
}
