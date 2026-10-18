package validator

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LoginMock struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=4,required"`
}

type UserMock struct {
	Name string   `json:"name" validate:"required,max=50"`
	Role RoleMock `json:"role"`
}

type RoleMock struct {
	Name string `json:"name" validate:"required,max=100"`
	Slug string `json:"slug" validate:"required,max=100"`
}

type PoolMock struct {
	Name      string         `json:"name" validate:"required,max=50"`
	Questions []QuestionMock `json:"questions"`
}

type QuestionMock struct {
	Title string `json:"title" validate:"required,max:50,trim"`
}

func TestValidatorWithInvalidData(t *testing.T) {
	jsonData := `{
		"email": "invalid-email",
		"password": "123"
	}`

	data := io.NopCloser(strings.NewReader(jsonData))

	credentials, errors := DecodeAndValidate[LoginMock](data)

	assert.Nil(t, credentials)
	assert.NotNil(t, errors)
}

func TestValidatorWithCompoundStructWithInvalidData(t *testing.T) {
	jsonData := `{
		"name": "",
		"role": {
			"name": "",
			"slug": ""
		}
	}`

	data := io.NopCloser(strings.NewReader(jsonData))

	user, errors := DecodeAndValidate[UserMock](data)

	assert.Nil(t, user)
	assert.NotNil(t, errors)
	assert.Len(t, errors, 2)
}

func TestValidatorWithInvalidStructSliceField(t *testing.T) {
	jsonData := `{
		"name": "",
		"questions": [
			{
				"title": "test title?"
			},
			{
				"title": ""
			}
		]
	}`

	data := io.NopCloser(strings.NewReader(jsonData))

	pool, errors := DecodeAndValidate[PoolMock](data)

	assert.Nil(t, pool)
	assert.NotNil(t, errors)
}

func TestValidatorWithValidStructSliceField(t *testing.T) {
	jsonData := `{
		"name": "Simple pool!",
		"questions": [
			{
				"title": "test title?"
			},
			{
				"title": "test title?"
			}
		]
	}`

	data := io.NopCloser(strings.NewReader(jsonData))

	pool, errors := DecodeAndValidate[PoolMock](data)

	assert.Nil(t, errors)
	assert.NotNil(t, pool)
}
