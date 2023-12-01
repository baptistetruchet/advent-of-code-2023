package utils

import "strings"

func Split(s string) []string {
	return strings.Split(s, "")
}

func SplitBy(c string) func(string) []string {
	return func(s string) []string {
		return strings.Split(s, c)
	}
}
