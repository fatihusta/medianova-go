package common

import (
	"fmt"
	"strconv"
	"strings"
)

// Number represents a JSON value that may be returned as an integer,
// a float, or a numeric string.
type Number string

func (n *Number) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		*n = ""
		return nil
	}
	if len(s) > 0 && (s[0] == '{' || s[0] == '[') {
		return fmt.Errorf("cannot unmarshal %s into Number", s)
	}
	*n = Number(strings.Trim(s, `"`))
	return nil
}

func (n Number) String() string {
	return string(n)
}

func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}
