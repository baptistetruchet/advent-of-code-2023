package utils

import "strconv"

func ParseInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func ParseFloat(s string) (f float64) {
	f, _ = strconv.ParseFloat(s, 64)
	return
}
