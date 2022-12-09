package validates

import (
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func PhoneNumber(fl validator.FieldLevel) bool {
	phone, ok := fl.Field().Interface().(string)
	if ok {
		locale := fl.Param()
		if len(locale) == 0 {
			locale = "CN"
		}
		num, err := phonenumbers.Parse(phone, locale)
		if err != nil {
			return false
		}
		return phonenumbers.IsValidNumberForRegion(num, locale)
	}
	return true
}
