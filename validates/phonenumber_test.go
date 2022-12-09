package validates

import (
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestPhoneNumber(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var test = struct {
				Phone string `json:"phone,omitempty" validate:"phone"`
			}{"18627977200"}
			en := en.New()
			uni := ut.New(en, en)

			// this is usually know or extracted from http 'Accept-Language' header
			// also see uni.FindTranslator(...)
			trans, _ := uni.GetTranslator("en")

			validate := validator.New()
			validate.RegisterValidation("phone", PhoneNumber)
			validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
				return ut.Add("phone", "{0} 手机号无效!", true) // see universal-translator for details
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("phone", fe.Field())

				return t
			})
			err := validate.Struct(test)
			if err != nil {
				errs := err.(validator.ValidationErrors)
				for _, e := range errs {
					// can translate each error one at a time.
					t.Errorf(e.Translate(trans))
				}
			} else {
				t.Logf("success validate ")
			}
		})
	}
}
