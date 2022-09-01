package validates

import (
	"github.com/dlclark/regexp2"

	"github.com/go-playground/validator/v10"
)

var _passwordPattern = regexp2.MustCompile("^(?![0-9]+$)(?![a-zA-Z]+$)(.){8,16}$", regexp2.Compiled)

func Password(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if ok {
		success, _ := _passwordPattern.MatchString(val)
		return success
	}
	return true
}
