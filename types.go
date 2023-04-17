package validator

type ModelInterface interface {
	SetFields()
	SetRules()
	Validate() bool
}

type Rule struct {
	FieldNames []string
	FType      string
	FRules     FieldRules
}

type Fields map[string]interface{}

type FieldRules struct {
	Max int
	Min int
}

type Model struct {
	Fields Fields
	Rules  []Rule
}

type ValidationErrors []error
