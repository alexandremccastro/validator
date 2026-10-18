package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

type ValidationError map[string]any

func DecodeAndValidate[T any](body io.ReadCloser) (*T, ValidationError) {
	var data T

	err := json.NewDecoder(body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
		return nil, map[string]any{
			"error": "failed to decode request",
		}
	}

	validation := walk(&data)

	if len(validation) > 0 {
		return nil, validation
	}

	return &data, nil
}

func walk[T any](data *T) ValidationError {
	return walkValue(reflect.ValueOf(data).Elem())
}

func walkValue(value reflect.Value) ValidationError {
	t := value.Type()

	errors := make(ValidationError)

	for index := range value.NumField() {
		structField := t.Field(index)
		fieldName := structField.Tag.Get("json")
		fieldValue := value.Field(index)

		switch fieldValue.Kind() {
		case reflect.Struct:
			validation := walkValue(fieldValue)
			if len(validation) > 0 {
				errors[fieldName] = validation
			}
		case reflect.Slice:
			if fieldValue.Index(0).Kind() == reflect.Struct {
				validation := walkSlice(fieldValue)

				if len(validation) > 0 {
					errors[fieldName] = validation
				}
			}
		default:
			rules := structField.Tag.Get("validate")
			field := NewField(fieldName, fieldValue, rules)

			validations := field.Validate()

			if len(validations) > 0 {
				errors[fieldName] = validations
			}
		}

	}

	return errors
}

func walkSlice(v reflect.Value) []ValidationError {
	var errors []ValidationError

	for index := range v.Len() {
		value := v.Index(index)

		validation := walkValue(value)

		if len(validation) > 0 {
			errors = append(errors, validation)
		}
	}

	return errors
}
