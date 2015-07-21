package int64utils

import (
	"strconv"
	"strings"
)

func Difference(a, b []int64) []int64 {
	out := make([]int64, 0)

	for _, v := range a {
		in := false

		for _, v2 := range b {
			if v2 == v {
				in = true
				break
			}
		}

		if !in {
			out = append(out, v)
		}
	}

	return out
}

func JoinString(a []int64, sep string) string {
	var s []string
	for _, i := range a {
		s = append(s, String(i))
	}
	return strings.Join(s, sep)
}

func String(i int64) string {
	return strconv.FormatInt(i, 10)
}
