package validates

import (
	"log"

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
		p, err := phonenumbers.Parse(phone, locale)
		log.Printf("%v", p)
		return err == nil
	}
	return true
}
