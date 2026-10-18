package validator

import (
	"reflect"
	"strings"
)

type Field struct {
	Name        string
	Value       reflect.Value
	Rules       []string
	Constraints map[string]string
}

func NewField(name string, value reflect.Value, validate string) *Field {
	validations := strings.Split(validate, ",")

	var rules []string
	constraints := make(map[string]string)

	for _, rule := range validations {
		if strings.Contains(rule, "=") {
			parts := strings.Split(rule, "=")
			rules = append(rules, parts[0])
			constraints[parts[0]] = parts[1]
		} else {
			rules = append(rules, rule)
		}
	}

	return &Field{
		Name:        name,
		Value:       value,
		Rules:       rules,
		Constraints: constraints,
	}
}

func (f *Field) Validate() []string {
	var errors []string

	for _, rule := range f.Rules {
		ruleFn, ok := RulesMap[rule]

		if !ok {
			continue
		}

		err := ruleFn(f)

		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	return errors
}
