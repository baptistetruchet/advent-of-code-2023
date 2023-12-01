package utils

import "strconv"

func ParseInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func ParseFloat(s string) (f float64) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return
}
