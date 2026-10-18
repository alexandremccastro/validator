package validator

import (
	"errors"
	"net/mail"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func IsEmail(f *Field) error {
	if f.Value.Kind() != reflect.String {
		return errors.New(ErrInvalidEmail)
	}

	email := f.Value.String()
	_, err := mail.ParseAddress(email)

	if err != nil {
		return errors.New(ErrInvalidEmail)
	}

	return nil
}

func IsRequired(f *Field) error {
	if f.Value.Kind() == reflect.String && f.Value.Len() == 0 {
		return errors.New(ErrIsRequired)
	}

	return nil
}

func Len(f *Field) error {
	expectedLen, err := strconv.ParseInt(f.Constraints["len"], 10, 64)

	if err != nil {
		return errors.New(ErrInvalidLength)
	}

	switch f.Value.Kind() {
	case reflect.String:
		if int64(f.Value.Len()) != expectedLen {
			return errors.New(ErrInvalidLength)
		}
	case reflect.Slice, reflect.Array:
		if int64(f.Value.Len()) != expectedLen {
			return errors.New(ErrInvalidLength)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := f.Value.Convert(reflect.TypeFor[int]()).Int()
		strValue := strconv.FormatInt(intValue, 10)

		if int64(len(strValue)) != expectedLen {
			return errors.New(ErrInvalidLength)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue := f.Value.Convert(reflect.TypeFor[uint]()).Uint()
		strValue := strconv.FormatUint(uintValue, 10)

		if int64(len(strValue)) != expectedLen {
			return errors.New(ErrInvalidLength)
		}
	case reflect.Float32, reflect.Float64:
		floatValue := f.Value.Convert(reflect.TypeFor[float64]()).Float()
		strValue := strconv.FormatFloat(floatValue, 'f', -1, 64)
		cleanValue := strings.ReplaceAll(strValue, ".", "")

		if int64(len(cleanValue)) != expectedLen {
			return errors.New(ErrInvalidLength)
		}

	default:
		return errors.New(ErrInvalidLength)
	}

	return nil
}

func Min(f *Field) error {
	minValue, err := strconv.ParseFloat(f.Constraints["min"], 64)

	if err != nil {
		return errors.New(ErrInvalidMin)
	}

	switch f.Value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := f.Value.Convert(reflect.TypeFor[int]()).Int()
		if float64(intValue) < minValue {
			return errors.New(ErrInvalidMin)
		}

	case reflect.String:
		if float64(len(f.Value.String())) < minValue {
			return errors.New(ErrInvalidMin)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue := f.Value.Convert(reflect.TypeFor[uint]()).Uint()
		if float64(uintValue) < minValue {
			return errors.New(ErrInvalidMin)
		}

	case reflect.Float32, reflect.Float64:
		floatValue := f.Value.Convert(reflect.TypeFor[float64]()).Float()
		if floatValue < minValue {
			return errors.New(ErrInvalidMin)
		}

	default:
		return errors.New(ErrInvalidMin)
	}

	return nil
}

func Max(f *Field) error {
	maxValue, err := strconv.ParseFloat(f.Constraints["max"], 64)

	if err != nil {
		return errors.New(ErrInvalidMax)
	}

	switch f.Value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := f.Value.Convert(reflect.TypeFor[int]()).Int()
		if float64(intValue) > maxValue {
			return errors.New(ErrInvalidMax)
		}

	case reflect.String:
		if float64(len(f.Value.String())) > maxValue {
			return errors.New(ErrInvalidMax)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue := f.Value.Convert(reflect.TypeFor[uint]()).Uint()
		if float64(uintValue) > maxValue {
			return errors.New(ErrInvalidMax)
		}

	case reflect.Float32, reflect.Float64:
		floatValue := f.Value.Convert(reflect.TypeFor[float64]()).Float()
		if floatValue > maxValue {
			return errors.New(ErrInvalidMax)
		}

	default:
		return errors.New(ErrInvalidMax)
	}

	return nil
}

func Trim(f *Field) error {
	if f.Value.Kind() == reflect.String && f.Value.CanSet() {
		f.Value.SetString(strings.Trim(f.Value.String(), " "))
	}

	return nil
}

func IsIn(f *Field) error {
	expectedValues := strings.Split(f.Constraints["in"], "|")

	if f.Value.Kind() != reflect.Slice {
		return errors.New(ErrInvalidIn)
	}

	element := f.Value.Index(0)

	switch element.Kind() {
	case reflect.String:
		for index := range f.Value.Len() {
			value := f.Value.Index(index)

			if !slices.Contains(expectedValues, value.String()) {
				return errors.New(ErrInvalidIn)
			}
		}
	}

	return nil
}
