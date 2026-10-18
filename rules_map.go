package validator

type Rule func(f *Field) error

var RulesMap = map[string]Rule{
	"email":    IsEmail,
	"required": IsRequired,
	"min":      Min,
	"max":      Max,
	"len":      Len,
	"trim":     Trim,
}
