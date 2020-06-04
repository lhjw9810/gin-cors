package types

import (
	"encoding/json"
	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Date time.Time
type Time time.Time
type DateTime time.Time

var jsonNull = []byte("null")

func (d *Date) UnmarshalJSON(data []byte) error {
	str, _ := strconv.Unquote(string(data[:]))
	_time, err := dateparse.ParseLocal(str)
	*d = Date(_time)
	return errors.WithStack(err)
}
func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Date().IsZero() {
		return jsonNull, nil
	}
	str := d.Date().Format("2006-01-02")
	return json.Marshal(str)
}

func (this Date) Date() time.Time {
	return (time.Time)(this)
}

func (this *Time) UnmarshalJSON(data []byte) error {
	str, _ := strconv.Unquote(string(data[:]))
	time, err := dateparse.ParseLocal("2006-01-02T" + str)
	*this = Time(time)
	return errors.WithStack(err)
}
func (this *Time) MarshalJSON() ([]byte, error) {
	if this.Time().IsZero() {
		return jsonNull, nil
	}
	str := this.Time().Format("15:04:05")
	return json.Marshal(str)
}

func (this Time) Time() time.Time {
	return (time.Time)(this)
}

func (this *DateTime) UnmarshalJSON(data []byte) error {
	str, _ := strconv.Unquote(string(data[:]))
	time, err := dateparse.ParseLocal(str)
	*this = DateTime(time)
	return errors.WithStack(err)
}
func (this *DateTime) MarshalJSON() ([]byte, error) {
	if this.DateTime().IsZero() {
		return jsonNull, nil
	}
	str := this.DateTime().Format("2006-01-02T15:04:05")
	return json.Marshal(str)
}

func (this DateTime) DateTime() time.Time {
	return (time.Time)(this)
}
