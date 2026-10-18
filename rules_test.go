package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmailWithInvalidEmail(t *testing.T) {
	v := reflect.ValueOf("invalid-email")
	field := NewField("test", v, "email")

	err := IsEmail(field)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidEmail)
}

func TestIsEmailWithValidEmail(t *testing.T) {
	v := reflect.ValueOf("test@example.com")
	field := NewField("test", v, "email")

	err := IsEmail(field)

	assert.Equal(t, err, nil)
}

func TestIsRequiredWithEmptyString(t *testing.T) {
	v := reflect.ValueOf("")

	field := NewField("test", v, "required")

	err := IsRequired(field)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrIsRequired)
}

func TestIsRequiredWithValidString(t *testing.T) {
	v := reflect.ValueOf("value")
	field := NewField("test", v, "required")

	err := IsRequired(field)
	assert.Equal(t, err, nil)
}

func TestLenWithInvalidLengthValues(t *testing.T) {
	intValue := reflect.ValueOf(12345)
	intField := NewField("test", intValue, "len=3")

	err := Len(intField)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidLength)

	floatValue := reflect.ValueOf(12.345)
	floatField := NewField("test", floatValue, "len=3")
	err = Len(floatField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidLength)

	uintValue := reflect.ValueOf(uint(12345))
	uintField := NewField("test", uintValue, "len=3")
	err = Len(uintField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidLength)

	stringValue := reflect.ValueOf("abcde")
	stringField := NewField("test", stringValue, "len=3")
	err = Len(stringField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidLength)

	sliceValue := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	sliceField := NewField("test", sliceValue, "len=3")
	err = Len(sliceField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidLength)
}

func TestLenWithvalidLengthValues(t *testing.T) {
	intValue := reflect.ValueOf(12345)
	intField := NewField("test", intValue, "len=5")

	err := Len(intField)

	assert.Equal(t, err, nil)

	floatValue := reflect.ValueOf(12.345)
	floatField := NewField("test", floatValue, "len=5")
	err = Len(floatField)
	assert.Equal(t, err, nil)

	uintValue := reflect.ValueOf(uint(12345))
	uintField := NewField("test", uintValue, "len=5")
	err = Len(uintField)
	assert.Equal(t, err, nil)

	stringValue := reflect.ValueOf("abcde")
	stringField := NewField("test", stringValue, "len=5")
	err = Len(stringField)
	assert.Equal(t, err, nil)

	sliceValue := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	sliceField := NewField("test", sliceValue, "len=5")
	err = Len(sliceField)
	assert.Equal(t, err, nil)
}

func TestMinWithInvalidValues(t *testing.T) {
	intValue := reflect.ValueOf(5)
	intField := NewField("test", intValue, "min=10")

	err := Min(intField)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidMin)

	floatValue := reflect.ValueOf(5.5)
	floatField := NewField("test", floatValue, "min=10.0")
	err = Min(floatField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidMin)

	uintValue := reflect.ValueOf(uint(5))
	uintField := NewField("test", uintValue, "min=10")
	err = Min(uintField)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), ErrInvalidMin)
}

func TestMinWithValidValues(t *testing.T) {
	intValue := reflect.ValueOf(15)
	intField := NewField("test", intValue, "min=10")

	err := Min(intField)

	assert.Equal(t, err, nil)

	floatValue := reflect.ValueOf(9.5)
	floatField := NewField("test", floatValue, "min=9.5")
	err = Min(floatField)
	assert.Equal(t, err, nil)

	uintValue := reflect.ValueOf(uint(15))
	uintField := NewField("test", uintValue, "min=5")
	err = Min(uintField)
	assert.Equal(t, err, nil)
}

func TestIsInWithInvalidContraints(t *testing.T) {
	data := reflect.ValueOf([]string{"copy"})
	field := NewField("role", data, "in=admin|reader")
	err := IsIn(field)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrInvalidIn)
}

func TestIsInWithValidConstraints(t *testing.T) {
	data := reflect.ValueOf([]string{"admin", "reader"})
	field := NewField("role", data, "in=admin|reader")
	err := IsIn(field)

	assert.Nil(t, err)
}
