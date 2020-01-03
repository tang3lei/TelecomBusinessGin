package tools

import (
	"regexp"
	"strconv"
)

var digitsRegexp = regexp.MustCompile(`[0-9]*`)

func StringGetInt(s string) int {
	l := digitsRegexp.FindStringSubmatch(s)
	r,_:= strconv.Atoi(l[0])
	return r
}
