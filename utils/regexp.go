package utils

import (
	"regexp"

	"github.com/samber/lo"
)

func FindAllStringGroup(r *regexp.Regexp, s string) [][]string {
	m := r.FindAllStringSubmatch(s, -1)
	return lo.Map(m, func(a []string, i int) []string { return a[1:] })
}
