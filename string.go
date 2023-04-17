package validator

import (
	"errors"
	"fmt"
	"reflect"
)

func NewModel(fields Fields, rules []Rule) *Model {
	return &Model{Rules: rules, Fields: fields}
}

func NewFieldRules(max int, min int) *FieldRules {
	return &FieldRules{
		Max: max,
		Min: min,
	}
}

func NewRule(fNames []string, fType string, fRules FieldRules) *Rule {
	return &Rule{
		FieldNames: fNames,
		FType:      fType,
		FRules:     fRules,
	}
}

func (m *Model) SetFields(fields Fields) {
	m.Fields = fields
}

func (m *Model) SetRules(rules []Rule) {
	m.Rules = rules
}

func (m *Model) Validate() (bool, ValidationErrors) {
	var validationErrors ValidationErrors

	for key, value := range m.Fields {
		rule, err := m.getRuleForField(key)

		if err != nil {
			validationErrors = append(validationErrors, err)
			continue
		}

		actualFType := reflect.TypeOf(value).String()

		if actualFType != rule.FType {
			err = errors.New(fmt.Sprintf("mismatched types. Expected: %s. Actual: %s", rule.FType, actualFType))
			validationErrors = append(validationErrors, err)
			continue
		}

		length := len(fmt.Sprintf("%v", value))

		if length < 1 {
			err = errors.New("value can not be empty")
			validationErrors = append(validationErrors, err)
			continue
		}

		if length > rule.FRules.Max {
			err = errors.New(fmt.Sprintf("value length can not be more than %d", rule.FRules.Max))
			validationErrors = append(validationErrors, err)
			continue
		}

		if length < rule.FRules.Min {
			err = errors.New(fmt.Sprintf("value length can not be less than %d", rule.FRules.Min))
			validationErrors = append(validationErrors, err)
			continue
		}
	}

	hasErrors := len(validationErrors) > 0

	fmt.Println(validationErrors)

	return !hasErrors, validationErrors
}

func (m *Model) getRuleForField(field string) (*Rule, error) {
	for _, rule := range m.Rules {
		for _, fieldName := range rule.FieldNames {
			if fieldName == field {
				return &rule, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("unknown field %s", field))
}

func ToStringErrors(errs ValidationErrors) string {
	var str string

	for _, err := range errs {
		str += "\n" + err.Error()
	}

	return str
}
