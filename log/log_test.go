package log

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func TestLogging(t *testing.T) {
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
	e := errors.WithStack(errors.New("no1"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Production(WithErrorRotate(100<<20, 30, 0, "log.log"))
			Errorf("%s", "this is debug formatter")
			Error(Fields{
				"url":    "https://",
				"method": "post",
			}, "this is debug field")
			WithErr(e, "with err")

		})
	}
}
