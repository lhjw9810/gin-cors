package validates

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestRegexPattern(t *testing.T) {
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
				Phone string `json:"phone,omitempty" validate:"password"`
			}{"18627977200"}

			validate := validator.New()
			validate.RegisterValidation("password", Password)
			err := validate.Struct(test)

			if err != nil {
				t.Errorf("fail on validate %+v", err)
			} else {
				t.Logf("success validate ")
			}
		})
	}
}
