package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

type Int64 int64

func (INT *Int64) UnmarshalJSON(data []byte) error {
	//data 字符串格式
	str, _ := strconv.Unquote(string(data[:])) // string(data[:])""123456"" 双引号格式 -> "123456"单引号格式
	//整形格式
	if str == "" {
		str = string(data[:])
	}
	i, err := strconv.ParseInt(str, 10, 64)
	*INT = Int64(i)
	return errors.WithStack(err)
}
func (INT *Int64) MarshalJSON() ([]byte, error) {
	str := strconv.FormatInt(int64(*INT), 10)
	return json.Marshal(str)
}
