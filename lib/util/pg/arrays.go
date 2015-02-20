package pg

import (
	"errors"
	"strconv"
	"strings"
)

type Integers []int

func (s *Integers) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	asString := string(asBytes)
	(*s) = strToIntSlice(asString)
	return nil
}

func strToIntSlice(s string) []int {
	r := strings.Trim(s, "{}")
	a := make([]int, 0, 10)
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.Atoi(t)
		a = append(a, i)
	}
	return a
}
